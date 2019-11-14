import re
from string import Template
from datetime import date
from unquietcode.tools.cfn_provider.models import ComplexType
from unquietcode.tools.cfn_provider.utils import snake_caps

from unquietcode.tools.cfn_provider.golang.code_model import (
    # GoPackage,
    SourceFile,
	GoParameter,
    GoFunction,
	GoMapLiteral,
	GoStructLiteral,
	ReturnExpression,
)
from unquietcode.tools.cfn_provider.golang.code_writer import write_file_to_string


# TODO make this dynamic
PROVIDER_VERSION = '0.0'

DEAD_LINE = '~x~x~x~x~x~x~x~'
MAX_PROPERTY_RECURSION = 5
PROJECT_URL = "https://github.com/UnquietCode/terraform-provider-cloudformation"


RESOURCE_ATTRIBUTE_TEMPLATE = Template(
"""
		"${name}": {
			Type: ${type},
			Elem: ${elem},
			Required: ${required},
			Optional: ${optional},
			ForceNew: ${force_replace},
			MaxItems: ${max_items},
			Set: ${set_function},
		},
"""[1:])


def _remove_dead_lines(content):
	lines = []

	for line in content.split('\n'):
		if DEAD_LINE in line:
			continue

		lines.append(line)

	content = '\n'.join(lines)
	content = content[:-1] if len(lines) > 1 else content
	return content


def _generate_attribute_struct(*, package_name, schema_name, attribute):
	attribute_type = attribute.type.type
	
	if isinstance(attribute_type, ComplexType):
		at_property_prefix = 'property'
		
		if attribute_type.package.full_path == 'cfn/misc':
			at_property_prefix = 'Property'
		
		if attribute_type == f"property{schema_name}":
			print("recursion")
		
		attribute_type = f'{at_property_prefix}{attribute_type.name}()'

	attribute_elem = attribute.type.element_type
	
	if attribute_elem is not None and not isinstance(attribute_elem, str):
		package_prefix = ""
		recursive = False
		
		if attribute_elem.package.name != package_name:
			package_prefix = attribute_elem.package.name + "."
		else:
			if attribute_elem.name == schema_name:
				recursive = True
	
		ae_property_prefix = 'property'
		
		if attribute_elem.package.full_path == 'cfn/misc':
			ae_property_prefix = 'Property'
		
		call = '()' if recursive is not True else '(strconv.Itoa(int(count + 1)))'
		attribute_elem = f'{package_prefix}{ae_property_prefix}{attribute_elem.name}{call}'
	
	name = snake_caps(attribute.name)
	
	struct = GoStructLiteral(
		type="",
		fields={
			"Type": attribute_type,
		}
	)
	
	def add_field(name, condition, value):
		if condition:
			struct.fields[name] = value
	
	add_field("Elem", attribute_elem, attribute_elem)
	add_field("Required", attribute.required is True, "true")
	add_field("Optional", attribute.required is not True and attribute.required is not None, "true")
	add_field("ForceNew", attribute.will_replace is True, "true")
	add_field("MaxItems", attribute.type.max_items is not None, attribute.type.max_items)
	add_field("Set", attribute.type.set_function, attribute.type.set_function)
	
	return name, struct


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
	

def _imports_stanza(imports):
	import_lines = [f'	"github.com/unquietcode/terraform-cfn-provider/{i}"' for i in sorted(imports)]
	import_lines = '\n'.join(import_lines)
	
	return import_lines or DEAD_LINE


def render_resource_template(*, cfn_version, imports, resource):
	import_lines = [ _import_prefix(i) for i in imports ]
	import_lines.extend([
		"github.com/hashicorp/terraform-plugin-sdk/helper/schema",
	])
	
	resource_name = resource.name
	package_name = resource.package.name
	attributes = resource.attributes.values()
	
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
			"Exists": f"resource{resource_name}Exists",
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
		fields={ name:value for name,value in attribute_structs },
	)
	
	# property_line="Required" if updatable else "Optional",
	
	functions = [
		GoFunction(
			name=f"Resource{resource_name}",
			parameters=[],
		    return_types=[
				"*schema.Resource"
		    ],
			body=[ReturnExpression(resource_struct)],
		),
		GoFunction(
			name=f"resource{resource_name}Exists",
			parameters=[
				GoParameter(name="data", type="*schema.ResourceData"),
				GoParameter(name="meta", type="interface{}"),
			],
			return_types=["bool", "error"],
			body=["return plugin.ResourceExists(data, meta)"],
		),
		GoFunction(
			name=f"resource{resource_name}Read",
			parameters=[
				GoParameter(name="data", type="*schema.ResourceData"),
				GoParameter(name="meta", type="interface{}"),
			],
			return_types=["error"],
			body=[f'return plugin.ResourceRead("{resource.cfn_type}", Resource{resource_name}(), data, meta)'],
		),
		GoFunction(
			name=f"resource{resource_name}Create",
			parameters=[
				GoParameter(name="data", type="*schema.ResourceData"),
				GoParameter(name="meta", type="interface{}"),
			],
			return_types=["error"],
			body=[f'return plugin.ResourceCreate("{resource.cfn_type}", Resource{resource_name}(), data, meta)'],
		),
		GoFunction(
			name=f"resource{resource_name}Update",
			parameters=[
				GoParameter(name="data", type="*schema.ResourceData"),
				GoParameter(name="meta", type="interface{}"),
			],
			return_types=["error"],
			body=[f'return plugin.ResourceUpdate("{resource.cfn_type}", Resource{resource_name}(), data, meta)'],
		),
		GoFunction(
			name=f"resource{resource_name}Delete",
			parameters=[
				GoParameter(name="data", type="*schema.ResourceData"),
				GoParameter(name="meta", type="interface{}"),
			],
			return_types=["error"],
			body=[f'return plugin.ResourceDelete("{resource.cfn_type}", data, meta)'],
		),
		GoFunction(
			name=f"resource{resource_name}CustomizeDiff",
			parameters=[
				GoParameter(name="data", type="*schema.ResourceDiff"),
				GoParameter(name="meta", type="interface{}"),
			],
			return_types=["error"],
			body=["return plugin.ResourceCustomizeDiff(data, meta)"],
		),
	]
	
	code_file = SourceFile(
	    header=_header_stanza(cfn_version, resource.documentation_link),
	    package_name=package_name,
	    imports=import_lines,
	    declarations=functions,
	)	
	
	rendered = write_file_to_string(code_file)
	rendered = _remove_dead_lines(rendered)
	return rendered



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


def render_property_template(*, cfn_version, package_name, property_name, attributes, imports, documentation_link):
	import_lines = [ _import_prefix(i) for i in imports ]
	import_lines.extend([
		"strconv",
		"github.com/hashicorp/terraform-plugin-sdk/helper/schema",
	])
	
	function_name = 'property' if property_name != 'Tag' else 'Property'
	function_name += property_name
	
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
		fields={ name:value for name,value in attribute_structs },
	)
	
	function_body = PROPERTY_FUNCTION_BODY.substitute(dict(
		max_recursion=MAX_PROPERTY_RECURSION,
	))
	
	functions = [
		GoFunction(
			name=function_name,
			parameters=[
				GoParameter(
					name="extras",
					type="string",
					varargs=True,
				)
			],
		    return_types=[
				"*schema.Resource"
		    ],
			body=[function_body, ReturnExpression(resource_struct)],
		)
	]
	
	code_file = SourceFile(
	    header=_header_stanza(cfn_version, documentation_link),
	    package_name=package_name,
	    imports=import_lines,
	    declarations=functions,
	)
	
	rendered = write_file_to_string(code_file)
	rendered = _remove_dead_lines(rendered)
	return rendered


def _resources_stanza(resources):
	resource_lines = [f'			"cfn_{r.service_name.lower()}_{snake_caps(r.name[len(r.service_name):])}": {r.package.name}.Resource{r.name}(),' for r in resources]
	resource_lines = '\n'.join(resource_lines)
	return resource_lines or DEAD_LINE


def _datasources_stanza(datasources):
	datasource_lines = [f'			"cfn_{r.service_name.lower()}_{snake_caps(r.name[len(r.service_name):])}": {d.package.name}.Datasource{d.name}(),' for d in datasources]
	datasource_lines = '\n'.join(datasource_lines)
	return datasource_lines or DEAD_LINE

	


PROVIDER_TEMPLATE = Template(
"""
${header}

package ${package}

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
${imports}
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: plugin.ProviderConfigure,
		Schema: map[string]*schema.Schema{
			"workdir": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "working directory on the filesystem",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cfn_template_data": misc.ResourceTemplateData(),
${resources}
		},
		DataSourcesMap: map[string]*schema.Resource{
${datasources}
		},
	}
}
"""[1:])


def render_provider_template(*, cfn_version, package_name, imports, datasources, resources):
	rendered = PROVIDER_TEMPLATE.substitute(dict(
		header=_header_stanza(cfn_version, None),
		package=package_name,
		imports=_imports_stanza(imports),
		resources=_resources_stanza(resources),
		datasources=_datasources_stanza(datasources),
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered

