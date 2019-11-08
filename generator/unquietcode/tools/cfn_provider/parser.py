from unquietcode.tools.cfn_provider.models import ResourceAttribute, Property, Resource, Package, Provider
from unquietcode.tools.cfn_provider.type_support import translate_cfn_resource_type#, translate_cfn_property_type


def handle_resource_property(property_name, property_data, schema_properties):
    attribute_type = translate_cfn_resource_type(property_data, schema_properties)
    required = property_data['Required']
    will_replace = property_data['UpdateType'] == 'Immutable'

    return ResourceAttribute(
        name=property_name,
        type=attribute_type,
        required=required,
        will_replace=will_replace,
    )

# TODO need to handle missing properties better

def handle_property(*, package, service, outer_name, inner_name, data):
    property_properties = data.get('Properties', {})
    attributes = {}

    for property_name, property_data in property_properties.items():
        attribute = handle_resource_property(property_name, property_data, package.properties)
        attributes[property_name] = attribute

    return Property(
        name=inner_name,
        package=package,
        resource_name=outer_name,
        attributes=attributes,
    )


def handle_resource(*, service, package, resource, cfn_type, resource_data):
    created_name = f"{service}{resource}"
    resource_properties = resource_data['Properties']
    attributes = {}

    for property_name, property_data in resource_properties.items():
        attribute = handle_resource_property(property_name, property_data, package.properties)
        attributes[property_name] = attribute

    resource = Resource(
        name=created_name,
        cfn_type=cfn_type,
        package=package,
        attributes=attributes,
    )

    return resource    
    

def _get_or_create_package(super_package, service):
    package_name = service.lower()

    if package_name not in super_package.subpackages:
        package = Package(name=package_name, parent=super_package)
        super_package.subpackages[package_name] = package
    else:
        package = super_package.subpackages[package_name]
    
    return package
    

def handle_special_prop(package, name, data):
    property = handle_property(
        package=package,
        service='',
        outer_name='',
        inner_name=name,
        data=data,
    )
    package.properties[name] = property


def handle_provider(data) -> Provider:
    super_package = Package(name='cfn')
    services_package = _get_or_create_package(super_package, 'services')
    misc_package = _get_or_create_package(super_package, 'misc')

    # do properties
    for property_name, property_data in data['PropertyTypes'].items():
        
        # special props
        if property_name in {'Tag'}:
            handle_special_prop(misc_package, property_name, property_data)
            continue
            
        parts = property_name.split('::')

        if parts[0] != "AWS":
            print(f"skipping non-aws property '{property_name}'")
            continue

        service = parts[1]
        subparts = parts[2].split('.')

        if len(subparts) != 2:
            raise Exception('too many name parts')

        outer_name, inner_name = subparts
        package = _get_or_create_package(services_package, service)
                    
        property = handle_property(
            package=package,
            service=service,
            outer_name=outer_name,
            inner_name=inner_name,
            data=property_data,
        )
        package.properties[property.name] = property
    
    # do resources
    for resource_name, resource_data in data['ResourceTypes'].items():
        parts = resource_name.split("::")

        if parts[0] != "AWS":
            print(f"skipping non-aws resource '{resource_name}'")
            continue

        service = parts[1]
        resource = parts[2]        
        package = _get_or_create_package(services_package, service)
        
        resource_object = handle_resource(
            service=service,
            package=package,
            resource=resource,
            cfn_type=resource_name,
            resource_data=resource_data,
        )
        package.resources[resource_object.name] = resource_object
    
    # metadata
    cfn_version = data['ResourceSpecificationVersion']
    
    return Provider(
        top_level_package=super_package,
        cfn_version=cfn_version,
    )