from unquietcode.tools.cfn_provider.models import AttributeType, CF_Type, TF_Type


TYPE_CONVERSION_MAP = {

    # primitive types
    CF_Type.String:    TF_Type.String,
    CF_Type.Long:      TF_Type.Int,
    CF_Type.Integer:   TF_Type.Int,
    CF_Type.Double:    TF_Type.Float,
    CF_Type.Boolean:   TF_Type.Bool,
    CF_Type.Timestamp: TF_Type.String,
    CF_Type.Json:      TF_Type.Map,
    
    # collection types
    CF_Type.Map:    TF_Type.Map,
    CF_Type.List:   TF_Type.List,
}


# def translate_cfn_property_type(property_name, property_data, schema_properties):
#     return property_name, None


def translate_cfn_resource_type(property_data, schema_properties) -> AttributeType:
    type = None
    elem_type = None
    max_items = None
    min_items = None
    set_function = None
    
    prop_ptype = property_data.get('PrimitiveType')

    # is a primitive type
    if prop_ptype is not None:
        type = TYPE_CONVERSION_MAP[CF_Type[prop_ptype]]

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
                elem_type = translate_cfn_complex_type(schema_properties, item_type)

        # is some other complex type
        else:
            type = TF_Type.List
            elem_type = translate_cfn_complex_type(schema_properties, prop_type)
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
    )


def translate_cfn_complex_type(schema_properties, name):
    if name is None:
        raise Exception('name is empty')
        
    if name in schema_properties:
        return schema_properties[name]
    
    # use deferred string
    return f'DEFERRED|{name}'