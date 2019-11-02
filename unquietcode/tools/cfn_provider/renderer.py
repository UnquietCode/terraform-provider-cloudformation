import os
import json
from .template import render_resource_template


def render_package(package, output_path):
    created_directory = f"{output_path}/{package.name}"
    os.makedirs(created_directory, exist_ok=True)

    for resource in package.resources.values():
        rendered = render_resource_template(
            package_name=package.name,
            resource_name=resource.name,
            attributes=resource.attributes,
        )

        with open(f"{created_directory}/{resource.file_name}", 'w+') as file_:
            file_.write(rendered)
        
    print(json.dumps(package.as_dict(), indent=2))
        
    for subpackage in package.subpackages.values():
        render_package(subpackage, created_directory)
