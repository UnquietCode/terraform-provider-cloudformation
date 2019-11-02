import os
import json

from .template import render_resource_template, render_property_template
from unquietcode.tools.cfn_provider.utils import snake_caps


def render_resource(package, resource, output_path):
    imports = set()
    
    for attr in resource.attributes:
        if attr.element in package.properties:
            prop = package.properties[attr.element]
            imports.add(prop.name)
    
    rendered = render_resource_template(
        package_name=f"cfn/{package.name}",
        resource_name=resource.name,
        attributes=resource.attributes,
        imports=imports,
    )

    with open(output_path, 'w+') as file_:
        file_.write(rendered)
    

def render_property(package, property, output_path):
    rendered = render_property_template(
        package_name=f"cfn/{package.name}",
        property_name=property.name,
        attributes=property.attributes,
    )

    with open(output_path, 'w+') as file_:
        file_.write(rendered)
    


def render_package(package, output_path):
    created_directory = f"{output_path}/{package.name}"
    os.makedirs(created_directory, exist_ok=True)

    for resource in package.resources.values():
        file_name = f"resource_{snake_caps(resource.name)}.go"
        file_path = f"{created_directory}/{file_name}"
        
        render_resource(package, resource, file_path)
    
    for property in package.properties.values():
        file_name = f"property_{snake_caps(property.name)}.go"
        file_path = f"{created_directory}/{file_name}"
        
        render_property(package, property, file_path)

    # print(json.dumps(package.as_dict(), indent=2))
        
    for subpackage in package.subpackages.values():
        render_package(subpackage, created_directory)
