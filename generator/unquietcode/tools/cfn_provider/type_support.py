from unquietcode.tools.cfn_provider.models import AttributeType, CF_Type, TF_Type, Validator


MISC_TYPES = {"Tag"}

TYPE_CONVERSION_MAP = {

    # primitive types
    CF_Type.String:    TF_Type.String,
    CF_Type.Long:      TF_Type.Int,
    CF_Type.Integer:   TF_Type.Int,
    CF_Type.Double:    TF_Type.Float,
    CF_Type.Boolean:   TF_Type.Bool,
    CF_Type.Timestamp: TF_Type.String,
    CF_Type.Json:      TF_Type.String,
    
    # collection types
    CF_Type.Map:    TF_Type.Map,
    CF_Type.List:   TF_Type.List,
}


def regex_validator(regex: str) -> Validator:
    return Validator(
        invocation=f'validation.StringMatch(regexp.MustCompile(`{regex}`), "must match pattern {regex}")',
        imports=[
            "regexp",
            "github.com/hashicorp/terraform-plugin-sdk/helper/validation",
        ]
    )


def simple_primitive(cf_type, validator=None) -> AttributeType:
    type = TYPE_CONVERSION_MAP[CF_Type[cf_type]]
    
    return AttributeType(
        type=type,
        element_type=None,
        min_items=None,
        max_items=None,
        set_function=None,
        validator_function=validator,
    )


def complex_type(type) -> AttributeType:
    return AttributeType(
        type=TF_Type.Set,
        element_type=type,
        min_items=None,
        max_items=None,
        set_function=None,
        validator_function=None,
    )


def translate_cfn_type(resource_name, property_data, schema_properties) -> AttributeType:
    type = None
    elem_type = None
    max_items = None
    min_items = None
    set_function = None
    validator_function = None
    
    prop_ptype = property_data.get('PrimitiveType')

    # is a primitive type
    if prop_ptype is not None:
        cf_type = CF_Type[prop_ptype]
        type = TYPE_CONVERSION_MAP[cf_type]
        
        # if it is JSON type, add validation
        if cf_type == CF_Type.Json:
            validator_function = Validator(
                invocation="validation.ValidateJsonString",
                imports=["github.com/hashicorp/terraform-plugin-sdk/helper/validation"],
            )

    # is a non-primitive type
    else:
        prop_type = property_data.get('Type')

        # is a collection type
        if prop_type in {CF_Type.List.value, CF_Type.Map.value}:
            prop_type = CF_Type[prop_type]
            allow_duplicates = property_data.get('DuplicatesAllowed')
            
            # translate further into Terraform collection types
            if allow_duplicates is False  and  prop_type is CF_Type.List:
                type = TF_Type.Set   
            else:
                type = TYPE_CONVERSION_MAP[prop_type]
            
            # look at element type
            pitem_type = property_data.get('PrimitiveItemType')
            
            # is a collection of primitives
            if pitem_type is not None:
                pitem_type = CF_Type[pitem_type]
                item_type = TYPE_CONVERSION_MAP[pitem_type]
                elem_type = item_type.schema_type

            # is a collection of some other type
            else:
                item_type = property_data.get('ItemType')
                elem_type = translate_cfn_complex_type(resource_name, item_type, schema_properties)

        # is some other complex type
        else:
            type = TF_Type.Set
            elem_type = translate_cfn_complex_type(resource_name, prop_type, schema_properties)
            max_items = 1
            
            # use required to determine if a minimum is needed
            required = property_data.get('Required')
            
            if required is True:
                min_items = 1

    # handle the set operations
    if type is TF_Type.Set:
        if elem_type == TF_Type.String.schema_type:
            set_function = 'schema.HashString'
        elif elem_type == TF_Type.Int.schema_type:
            set_function = 'schema.HashInt'
    
    return AttributeType(
        type=type,
        element_type=elem_type,
        min_items=min_items,
        max_items=max_items,
        set_function=set_function,
        validator_function=validator_function,
    )


def translate_cfn_complex_type(resource_name, name, schema_properties):
    if name is None:
        raise Exception('name is empty')
    
    if name in schema_properties:
        return schema_properties[name]
    
    # use deferred string
    
    # unfortunate special handling
    if name in MISC_TYPES:
        return f'DEFERRED|{name}'
    
    # normal case
    else:
        return f'DEFERRED|{resource_name}{name}'