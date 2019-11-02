

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


def translate_cfn_type(property_data):
    ptype = property_data.get('PrimitiveType')
    type = None
    elemType = None

    # is a primitive type
    if ptype is not None:
        type = PRIMITIVE_TYPES_MAP[ptype]

    # is a non-primitive type
    else:
        ptype = property_data.get('Type')

        # is a collection type
        if ptype == 'List' or ptype == 'Map':
            type = COLLECTION_TYPES_MAP[ptype]
            itemType = property_data.get('PrimitiveItemType')

            # is a collection of primitives
            if itemType is not None:
                itemType = PRIMITIVE_TYPES_MAP[itemType]
                itemType = "&schema.Schema{Type: "+itemType+"}"

            # is a collection of some other type
            else:
                itemType = property_data.get('ItemType')
                print("need handling for type "+itemType)
                itemType = 'schema.TypeString'

            elemType = itemType

        # is some other type
        else:
            print("need handling for type "+ptype)
            type = 'schema.TypeString'

    return type, elemType



def translate_cfn_complex_type(name, propData):
    pass
