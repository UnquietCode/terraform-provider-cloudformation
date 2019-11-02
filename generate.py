import os
import json

from template import render_resource_template
from resource_attribute import ResourceAttribute
from unquietcode.tools.cfn_provider.type_support import translate_cfn_type
from unquietcode.tools.cfn_provider.utils import snake_caps

THIS_DIR = os.path.dirname(os.path.realpath(__file__))
SPEC_FILE = f"{THIS_DIR}/CloudFormationResourceSpecification-us-east-1.json"
OUT_DIR = f"{THIS_DIR}/out"


def process_property(property_name, property_data):
    type, elemType = translate_cfn_type(property_data)
    required = property_data['Required']
    will_replace = property_data['UpdateType'] == 'Immutable'

    return ResourceAttribute(
        name=property_name,
        type=type,
        element=elemType,
        required=required,
        will_replace=will_replace,
    )


def process_resource(resource_data, created_package, created_name):
    resource_properties = resource_data['Properties']
    attributes = []

    for property_name, property_data in resource_properties.items():
        attribute = process_property(property_name, property_data)
        attributes.append(attribute)

    rendered = render_resource_template(
        package_name=created_package,
        resource_name=created_name,
        attributes=attributes,
    )

    return rendered



def handle_property(service, outer_name, inner_name, data):
    print(f"-------- {outer_name} -------- {inner_name}")


def handle_resource(service, resource, data):
    created_package = f"cfn/{service.lower()}"
    created_name = f"{service}{resource}"
    created_file = f"resource_{snake_caps(created_name)}.go"

    created_directory = f"{OUT_DIR}/cfn/{service.lower()}"
    os.makedirs(created_directory, exist_ok=True)

    rendered = process_resource(data, created_package, created_name)

    with open(f"{created_directory}/{created_file}", 'w+') as file_:
        file_.write(rendered)



def main():

    with open(SPEC_FILE, 'r+') as file_:
        data = json.load(file_)

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
        handle_property(service, outer_name, inner_name, property_data)

    for resource_name, resource_data in data['ResourceTypes'].items():
        parts = resource_name.split("::")

        if parts[0] != "AWS":
            print(f"skipping non-aws resource '{resource_name}'")
            continue

        service = parts[1]
        resource = parts[2]
        handle_resource(service, resource, resource_data)


if __name__ == '__main__':
    main()
