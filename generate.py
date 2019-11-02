import os
import json

from unquietcode.tools.cfn_provider.renderer import render_package
from unquietcode.tools.cfn_provider.models import ResourceAttribute, Package, Property, Resource
from unquietcode.tools.cfn_provider.type_support import translate_cfn_resource_type, translate_cfn_property_type
from unquietcode.tools.cfn_provider.utils import snake_caps

THIS_DIR = os.path.dirname(os.path.realpath(__file__))
SPEC_FILE = f"{THIS_DIR}/CloudFormationResourceSpecification-us-east-1.json"
OUT_DIR = f"{THIS_DIR}/out"


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



def main():
    package = generate_package()
    render_package(package, OUT_DIR)
    
    
def generate_package():
    
    with open(SPEC_FILE, 'r+') as file_:
        data = json.load(file_)

    super_package = Package(name='cfn')
    
    # do properties
    for property_name, property_data in data['PropertyTypes'].items():
        parts = property_name.split('::')

        if parts[0] != "AWS":
            print(f"skipping non-aws property '{property_name}'")
            continue

        service = parts[1]
        subparts = parts[2].split('.')

        if len(subparts) != 2:
            raise Exception('too many name parts')

        outer_name, inner_name = subparts
        package_name = service.lower()

        if package_name not in super_package.subpackages:
            package = Package(name=package_name)
            super_package.subpackages[package_name] = package
        else:
            package = super_package.subpackages[package_name]
                    
        # super_package.properties[service] = super_package.properties.get(service, {})
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
        package_name = service.lower()
        
        if package_name not in super_package.subpackages:
            package = Package(name=package_name)
            super_package.subpackages[package_name] = package
        else:
            package = super_package.subpackages[package_name]
        
        resource_object = handle_resource(service, package, resource, resource_data)
        package.resources[resource_object.name] = resource_object
        
    return super_package


if __name__ == '__main__':
    main()
