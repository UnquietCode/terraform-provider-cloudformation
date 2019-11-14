from unquietcode.tools.cfn_provider.models import ResourceAttribute, Property, Resource, Package, Provider
from unquietcode.tools.cfn_provider.type_support import translate_cfn_type, simple_primitive

RESERVED_PROPERTY_ATTRIBUTES = {'Count', 'Provider'}
RESERVED_RESOURCE_ATTRIBUTES = {'Id'}
RESERVED_RESOURCE_ATTRIBUTES.update(RESERVED_PROPERTY_ATTRIBUTES)


def handle_resource_property(resource_name, property_name, property_data, schema_properties, computed, type):
    attribute_type = translate_cfn_type(resource_name, property_data, schema_properties)
    reserved = RESERVED_PROPERTY_ATTRIBUTES if type == 'property' else RESERVED_RESOURCE_ATTRIBUTES
    
    if property_name in reserved:
        property_name = f"The{property_name}"

    return ResourceAttribute(
        name=property_name,
        type=attribute_type,
        required=property_data['Required'],
        computed=False,#computed,
        will_replace=False, #property_data['UpdateType'] == 'Immutable',
        documentation_link=property_data['Documentation'],
    )


def handle_resource_attribute(resource_name, documentation_link, attribute_name, attribute_data):
    attribute_type = translate_cfn_type(resource_name, attribute_data, {})
    
    if attribute_name in RESERVED_RESOURCE_ATTRIBUTES:
        attribute_name = f"The{attribute_name}"
    
    # replace nested with flat for now
    attribute_name = attribute_name.replace('.', '')
    
    return ResourceAttribute(
        name=attribute_name,
        type=attribute_type,
        computed=True,
        required=None,
        will_replace=None,
        documentation_link=documentation_link,
    )


# TODO need to handle missing properties better

def handle_property(*, package, service, outer_name, inner_name, data):
    property_properties = data.get('Properties', {})
    attributes = {}
    
    for property_name, property_data in property_properties.items():
        attribute = handle_resource_property(outer_name, property_name, property_data, package.properties, False, 'property')
        attributes[attribute.name] = attribute

    return Property(
        name=f"{outer_name}{inner_name}",
        package=package,
        documentation_link=data.get('Documentation'),
        attributes=attributes,
    )


def handle_resource(*, service, package, resource_name, cfn_type, resource_data):
    
    # resource properties (clobbers attributes with same name)
    resource_properties = resource_data['Properties']
    attributes = {}
    
    for property_name, property_data in resource_properties.items():
        
        # computable if an attribute exists and it is not required
        computed = property_name in resource_properties \
               and property_data.get('Required') is not True
        
        attribute = handle_resource_property(
            resource_name=resource_name,
            property_name=property_name,
            property_data=property_data,
            schema_properties=package.properties,
            computed=computed,
            type='resource',
        )
        attributes[attribute.name] = attribute
    
    # resource attributes
    # resource_attributes = resource_data.get("Attributes", {})
    # attributes = {}
    # 
    # for attribute_name, attribute_data in resource_attributes.items():
    #     attribute = handle_resource_attribute(
    #         resource_name=resource_name,
    #         documentation_link=attribute_data.get('Documentation'),
    #         attribute_name=attribute_name,
    #         attribute_data=attribute_data,
    #     )
    #     attributes[attribute.name] = attribute
    
    resource = Resource(
        name=f"{service}{resource_name}",
        service_name=service,
        resource_name=resource_name,
        documentation_link=resource_data['Documentation'],
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
            resource_name=resource,
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