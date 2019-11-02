from unquietcode.tools.cfn_provider.models import ResourceAttribute, Property, Resource
from unquietcode.tools.cfn_provider.type_support import translate_cfn_resource_type, translate_cfn_property_type
from unquietcode.tools.cfn_provider.utils import snake_caps


def process_resource_property(property_name, property_data, schema_properties):
    type, elemType = translate_cfn_resource_type(property_data, schema_properties)
    required = property_data['Required']
    will_replace = property_data['UpdateType'] == 'Immutable'

    return ResourceAttribute(
        name=property_name,
        type=type,
        element=elemType,
        required=required,
        will_replace=will_replace,
    )


def process_resource(*, resource_data, package, service, created_name, created_file):
    resource_properties = resource_data['Properties']
    attributes = []

    for property_name, property_data in resource_properties.items():
        attribute = process_resource_property(property_name, property_data, package.properties)
        attributes.append(attribute)

    resource = Resource(
        name=created_name,
        file_name=created_file,
        attributes=attributes,
    )

    return resource



def handle_property(*, package, service, outer_name, inner_name, data):
    # print(f"-------- {outer_name} -------- {inner_name}")
    type, elemType = translate_cfn_property_type(inner_name, data, package.properties)
    
    return Property(
        name=inner_name,
        package_name=service,
        resource_name=outer_name,
        type=type,
    )


def handle_resource(service, package, resource, data):    
    created_name = f"{service}{resource}"
    created_file = f"resource_{snake_caps(created_name)}.go"
    
    resource = process_resource(
        resource_data=data,
        package=package,
        service=service,
        created_name=created_name,
        created_file=created_file,
    )
    
    return resource