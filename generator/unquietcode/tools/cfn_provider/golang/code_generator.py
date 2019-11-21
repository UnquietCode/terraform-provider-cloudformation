import re
from datetime import date
from string import Template

from unquietcode.tools.cfn_provider.models import Property, TF_Type, GetAttr
from unquietcode.tools.cfn_provider.utils import snake_caps
from unquietcode.tools.cfn_provider.golang.code_model import (
    SourceFile,
    GoParameter,
    GoFunction,
    GoMapLiteral,
    GoStructLiteral,
    ReturnExpression,
    LiteralExpression,
    AssignmentExpression,
    ConstantExpression,
    BlankLines,
)
from unquietcode.tools.cfn_provider.golang.code_writer import write_file_to_string


# TODO make this dynamic
PROVIDER_VERSION = '0.0'

MAX_PROPERTY_RECURSION = 5
PROJECT_URL = "https://github.com/UnquietCode/terraform-provider-cloudformation"


def _generate_attribute_struct(*, package_name, schema_name, attribute):
    attribute_type = attribute.type.type
    name = snake_caps(attribute.name)
    is_gett_attr = False
    
    if isinstance(attribute_type, Property):
        at_property_prefix = 'property'
        at_name = attribute_type.name
        
        if attribute_type.package.full_path == 'cfn/misc':
            at_property_prefix = 'Property'
            
            if at_name == 'Tag':
                return (name, attribute.name), LiteralExpression("misc.PropertyTags()")
        
        if attribute_type == f"property{schema_name}":
            print("recursion")
        
        attribute_type = f'{at_property_prefix}{at_name}()'
    
    attribute_elem = attribute.type.element_type
    
    if attribute_elem is not None and not isinstance(attribute_elem, str):
        package_prefix = ""
        recursive = False
        
        if attribute_elem.package.name != package_name:
            package_prefix = attribute_elem.package.name + "."
        else:
            if attribute_elem.name == schema_name:
                recursive = True
        
        if isinstance(attribute_elem, GetAttr):
            attribute_elem = f'attributes{attribute_elem.resource_name}()'
            is_gett_attr = True
        
        else:
            ae_property_prefix = 'property'
            ae_name = attribute_elem.name
            
            if attribute_elem.package.full_path == 'cfn/misc':
                ae_property_prefix = 'Property'
                
                if ae_name == 'Tag':
                    return (name, attribute.name), LiteralExpression("misc.PropertyTags()")
            
            call = '()' if recursive is not True else '(strconv.Itoa(int(count + 1)))'
            attribute_elem = f'{package_prefix}{ae_property_prefix}{ae_name}{call}'
        
    struct = GoStructLiteral(
        type="",
        fields={
            "Type": attribute_type,
        }
    )
    
    def add_field(name, condition, value):
        if condition:
            if callable(value):
                value = value()
            
            struct.fields[name] = value
    
    add_field("Elem", attribute_elem, attribute_elem)
    add_field("Required", attribute.required is True, "true")
    add_field("Optional", attribute.required is not True and attribute.required is not None, "true")
    add_field("ForceNew", attribute.will_replace is True, "true")
    add_field("MaxItems", attribute.type.max_items is not None, attribute.type.max_items)
    add_field("Set", attribute.type.set_function, attribute.type.set_function)
    add_field("Computed", attribute.computed is True, "true")
    add_field("ValidateFunc", attribute.type.validator_function, lambda: attribute.type.validator_function.invocation)
        
    if is_gett_attr is True:
        add_field("StateFunc", True, 'func(v interface{}) string { return "" }')
    
    return (name, attribute.name), struct


HEADER_TEMPLATE = Template(
"""
// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on ${date}, using version ${provider_version} of the cfn terraform provider,
// and version ${schema_version} of the CloudFormation resource specification.
//
// For more information, visit:
//   ${documentation_link}
"""[1:-1])


def _header_stanza(cfn_version, documentation_link):
    return HEADER_TEMPLATE.substitute(dict(
        date=date.today().strftime("%d-%m-%Y"),
        provider_version=PROVIDER_VERSION,
        schema_version=cfn_version,
        documentation_link=documentation_link or PROJECT_URL,
    ))


def _import_prefix(i):
    return f"github.com/unquietcode/terraform-cfn-provider/{i}"
    

def _property_lookup(entity_name, attribute_structs):
    property_lookup = {
        names[0]: LiteralExpression(f'"{names[1]}"')
        for names,value in attribute_structs
    }
    
    # remove certain things which were added
    property_lookup.pop("logical_id", None)
    map_name = f"{entity_name[0].lower()}{entity_name[1:]}Properties"
    
    expression = AssignmentExpression(
        identifier=map_name,
        type="map[string]string",
        expression=GoMapLiteral(
            key_type="string",
            value_type="string",
            fields=property_lookup,
        )
    )
    
    return map_name, expression


def generate_resource_model(*, cfn_version, imports, resource):
    resource_name = resource.name
    package_name = resource.package.name
    attributes = resource.attributes.values()

    import_lines = [ _import_prefix(i) for i in imports ]
    import_lines.extend([
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema",
    ])
    
    for attribute in attributes:
        if attribute.type.validator_function:
            import_lines.extend(attribute.type.validator_function.imports)
    
    import_lines = list(set(import_lines))
    
    attribute_structs = [
        _generate_attribute_struct(package_name=package_name, schema_name=None, attribute=attribute)
        for attribute in attributes
    ]
    
    # Terraform requires that we check for the case of non-updatable resources
    updatable = False
    
    # at least one attribute will not force a replacement
    for attribute in attributes:
        if attribute.will_replace is not True and attribute.computed is not True:
            updatable = True
            break
    
    resource_struct = GoStructLiteral(
        type="&schema.Resource",
        fields={
            # "Exists": f"resource{resource_name}Exists",
            "Read": f"resource{resource_name}Read",
            "Create": f"resource{resource_name}Create",
            "Update": f"resource{resource_name}Update",
            "Delete": f"resource{resource_name}Delete",
            "CustomizeDiff": f"resource{resource_name}CustomizeDiff",
        }
    )
    
    if not updatable:
        del resource_struct.fields["Update"]
    
    resource_struct.fields["Schema"] = GoMapLiteral(
        key_type="string",
        value_type="*schema.Schema",
        fields={ names[0]:value for names,value in attribute_structs },
    )
    
    type_const = f"{resource_name[0].lower()}{resource_name[1:]}Type"
    # property_map_name, property_map = _property_lookup(resource_name, attribute_structs)
    
    declarations = [
        ConstantExpression(
            identifier=type_const,
            type="string",
            expression=LiteralExpression(f'"{resource.cfn_type}"'),
        ),
        BlankLines(count=1),
        
        # property_map,
        
        GoFunction(
            name=f"Resource{resource_name}",
            parameters=[],
            return_types=[
                "*schema.Resource"
            ],
            body=[ReturnExpression(resource_struct)],
        ),
        # GoFunction(
        #     name=f"resource{resource_name}Exists",
        #     parameters=[
        #         GoParameter(name="data", type="*schema.ResourceData"),
        #         GoParameter(name="meta", type="interface{}"),
        #     ],
        #     return_types=["bool", "error"],
        #     body=["return plugin.ResourceExists(data, meta)"],
        # ),
        GoFunction(
            name=f"resource{resource_name}Read",
            parameters=[
                GoParameter(name="data", type="*schema.ResourceData"),
                GoParameter(name="meta", type="interface{}"),
            ],
            return_types=["error"],
            body=[f'return plugin.ResourceRead({type_const}, Resource{resource_name}(), data, meta)'],
        ),
        GoFunction(
            name=f"resource{resource_name}Create",
            parameters=[
                GoParameter(name="data", type="*schema.ResourceData"),
                GoParameter(name="meta", type="interface{}"),
            ],
            return_types=["error"],
            body=[f'return plugin.ResourceCreate({type_const}, Resource{resource_name}(), data, meta)'],
        ),
        GoFunction(
            name=f"resource{resource_name}Update",
            parameters=[
                GoParameter(name="data", type="*schema.ResourceData"),
                GoParameter(name="meta", type="interface{}"),
            ],
            return_types=["error"],
            body=[f'return plugin.ResourceUpdate({type_const}, Resource{resource_name}(), data, meta)'],
        ),
        GoFunction(
            name=f"resource{resource_name}Delete",
            parameters=[
                GoParameter(name="data", type="*schema.ResourceData"),
                GoParameter(name="meta", type="interface{}"),
            ],
            return_types=["error"],
            body=[f'return plugin.ResourceDelete({type_const}, data, meta)'],
        ),
        GoFunction(
            name=f"resource{resource_name}CustomizeDiff",
            parameters=[
                GoParameter(name="data", type="*schema.ResourceDiff"),
                GoParameter(name="meta", type="interface{}"),
            ],
            return_types=["error"],
            body=[f'return plugin.ResourceCustomizeDiff({type_const}, data, meta)'],
        ),
    ]
    
    code_file = SourceFile(
        header=_header_stanza(cfn_version, resource.documentation_link),
        package_name=package_name,
        imports=import_lines,
        declarations=declarations,
    )    
    
    return write_file_to_string(code_file)


PROPERTY_FUNCTION_BODY = Template(
"""
var count int64 = 0

if len(extras) > 0 {
    if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
        count = i
    }
}

if count >= ${max_recursion} {
    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
}

"""[1:])


def generate_property_model(*, cfn_version, package_name, property_name, attributes, imports, documentation_link):
    import_lines = [ _import_prefix(i) for i in imports ]
    import_lines.extend([
        "strconv",
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema",
    ])
    
    for attribute in attributes:
        if attribute.type.validator_function:
            import_lines.extend(attribute.type.validator_function.imports)    
    
    import_lines = list(set(import_lines))
    
    attribute_structs = [
        _generate_attribute_struct(package_name=package_name, schema_name=property_name, attribute=attribute)
        for attribute in attributes
    ]
    
    resource_struct = GoStructLiteral(
        type="&schema.Resource",
    )
    
    resource_struct.fields["Schema"] = GoMapLiteral(
        key_type="string",
        value_type="*schema.Schema",
        fields={ names[0]:value for names,value in attribute_structs },
    )
        
    function_name = 'property' if property_name != 'Tag' else 'Property'
    function_name += property_name
    
    function_body = PROPERTY_FUNCTION_BODY.substitute(dict(
        max_recursion=MAX_PROPERTY_RECURSION,
    ))
    
    # property_map_name, property_map = _property_lookup(property_name, attribute_structs)
    
    declarations = [
        # property_map,
        
        GoFunction(
            name=function_name,
            parameters=[
                GoParameter(
                    name="extras",
                    type="string",
                    varargs=True,
                )
            ],
            return_types=["*schema.Resource"],
            body=[function_body, ReturnExpression(resource_struct)],
        )
    ]
    
    code_file = SourceFile(
        header=_header_stanza(cfn_version, documentation_link),
        package_name=package_name,
        imports=import_lines,
        declarations=declarations,
    )
    
    return write_file_to_string(code_file)


def generate_getattr_model(*, cfn_version, package_name, get_attr, documentation_link):
    import_lines = [
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema",
    ]
    
    for name, attribute in get_attr.attributes.items():
        if attribute.type.validator_function:
            import_lines.extend(attribute.type.validator_function.imports)    
    
    import_lines = list(set(import_lines))
    
    attribute_structs = [
        _generate_attribute_struct(package_name=package_name, schema_name=None, attribute=attribute)
        for attribute in get_attr.attributes.values()
    ]
    
    resource_struct = GoStructLiteral(
        type="&schema.Resource",
    )
    
    resource_struct.fields["Schema"] = GoMapLiteral(
        key_type="string",
        value_type="*schema.Schema",
        fields={ names[0]:value for names,value in attribute_structs },
    )
        
    function_name = f'attributes{get_attr.resource_name}'
        
    declarations = [
        GoFunction(
            name=function_name,
            parameters=[],
            return_types=["*schema.Resource"],
            body=[ReturnExpression(resource_struct)],
        )
    ]
    
    code_file = SourceFile(
        header=_header_stanza(cfn_version, documentation_link),
        package_name=package_name,
        imports=import_lines,
        declarations=declarations,
    )
    
    return write_file_to_string(code_file)


def _resource_line(resource):
    r = resource
    
    name = f'cfn_{r.service_name.lower()}_{snake_caps(r.name[len(r.service_name):])}'
    value = f'{r.package.name}.Resource{r.name}()'
    return name, LiteralExpression(expression=value)


def _datasources_line(datasource):
    d = datasource
    
    name = f'cfn_{r.service_name.lower()}_{snake_caps(r.name[len(r.service_name):])}'
    value = f'{d.package.name}.Datasource{d.name}(),'
    return name, LiteralExpression(expression=value)
    

def generate_provider_model(*, cfn_version, package_name, imports, datasources, resources):
    import_lines = [ _import_prefix(i) for i in imports ]
    import_lines.extend([
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema",
        "github.com/hashicorp/terraform-plugin-sdk/terraform",
        "github.com/unquietcode/terraform-cfn-provider/plugin",
        "github.com/unquietcode/terraform-cfn-provider/cfn/misc",
    ])
    
    resource_fields = [ _resource_line(_) for _ in resources ]
    resource_fields = { _[0]:_[1] for _ in resource_fields }
    resource_fields["cfn_template_data"] =  LiteralExpression(expression="misc.ResourceTemplateData()")

    datasource_fields = [ _datasources_line(_) for _ in datasources ]
    datasource_fields = { _[0]:_[1] for _ in datasource_fields }
    
    provider_struct = GoStructLiteral(
        type="&schema.Provider",
        fields={
            "ConfigureFunc": "plugin.ProviderConfigure",
            "Schema": GoMapLiteral(
                key_type="string",
                value_type="*schema.Schema",
                fields={
                    "workdir": GoStructLiteral(
                        type="",
                        fields={
                            "Type": "schema.TypeString",
                            "Required": "true",
                            "Description": '"working directory on the filesystem"',
                        }
                    )
                }
            ),
            "ResourcesMap": GoMapLiteral(
                key_type="string",
                value_type="*schema.Resource",
                fields=resource_fields,
            ),
            "DataSourcesMap": GoMapLiteral(
                key_type="string",
                value_type="*schema.Resource",
                fields=datasource_fields,
            )
        },
    )
    
    functions = [
        GoFunction(
            name="Provider",
            parameters=[],
            return_types=["terraform.ResourceProvider"],
            body=[ReturnExpression(provider_struct)],
        )
    ]    
    
    code_file = SourceFile(
        header=_header_stanza(cfn_version, None),
        package_name=package_name,
        imports=import_lines,
        declarations=functions,
    )
    
    return write_file_to_string(code_file)