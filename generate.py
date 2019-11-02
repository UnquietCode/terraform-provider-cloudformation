import os
import json

from unquietcode.tools.cfn_provider.models import Package
from unquietcode.tools.cfn_provider.parser import handle_property, handle_resource
from unquietcode.tools.cfn_provider.renderer import render_package

THIS_DIR = os.path.dirname(os.path.realpath(__file__))
SPEC_FILE = f"{THIS_DIR}/CloudFormationResourceSpecification-us-east-1.json"
OUT_DIR = f"{THIS_DIR}/out"



def main():
    package = generate_package()
    render_package(package, OUT_DIR)


def get_or_create_package(super_package, service):
    package_name = service.lower()

    if package_name not in super_package.subpackages:
        package = Package(name=package_name)
        super_package.subpackages[package_name] = package
    else:
        package = super_package.subpackages[package_name]
    
    return package
    

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
        package = get_or_create_package(super_package, service)
                    
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
        package = get_or_create_package(super_package, service)
        
        resource_object = handle_resource(service, package, resource, resource_data)
        package.resources[resource_object.name] = resource_object
    
        
    return super_package


if __name__ == '__main__':
    main()
