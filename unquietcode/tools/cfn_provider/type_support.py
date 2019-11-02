

PRIMITIVE_TYPES_MAP = {
    'String': 'schema.TypeString',
    'Long': 'schema.TypeInt',
    'Integer': 'schema.TypeInt',
    'Double': 'schema.TypeFloat',
    'Boolean': 'schema.TypeBool',
    'Timestamp': 'schema.TypeString',
    'Json': 'schema.TypeMap',
}

COLLECTION_TYPES_MAP = {
    'Map': 'schema.TypeMap',
    'List': 'schema.TypeList',
}


def translate_cfn_property_type(property_name, property_data, schema_properties):
    return property_name, None

def translate_cfn_resource_type(property_data, schema_properties):
    prop_ptype = property_data.get('PrimitiveType')
    type = None
    elemType = None

    # is a primitive type
    if prop_ptype is not None:
        type = PRIMITIVE_TYPES_MAP[prop_ptype]

    # is a non-primitive type
    else:
        prop_type = property_data.get('Type')

        # is a collection type
        if prop_type == 'List' or prop_type == 'Map':
            type = COLLECTION_TYPES_MAP[prop_type]
            pitem_type = property_data.get('PrimitiveItemType')

            # is a collection of primitives
            if pitem_type is not None:
                item_type = PRIMITIVE_TYPES_MAP[pitem_type]
                elemType = "&schema.Schema{Type: "+item_type+"}"

            # is a collection of some other type
            else:
                item_type = property_data.get('ItemType')
                
                # if item_type is not None:
                elemType = item_type
                # else:
                    # elemType = 'schema.TypeString'

        # is some other type
        else:
            type = translate_cfn_complex_type(schema_properties, prop_type, property_data)

    return type, elemType



def translate_cfn_complex_type(schema_properties, name, propData):
    if name is None:
        raise Exception('name is empty')
        
    if name in schema_properties:
        print(f"found {name}")
    else:
        print(f"need handling for type {name}")
    
    return 'schema.TypeString'
