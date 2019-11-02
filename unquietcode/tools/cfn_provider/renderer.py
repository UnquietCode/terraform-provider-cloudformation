import os
import json

from .template import render_resource_template, render_property_template
from unquietcode.tools.cfn_provider.utils import snake_caps


def _extract_imports(attributes):
    imports = set()
    
    for attr in attributes:
        if not isinstance(attr.type, str):
            imports.add(f'{attr.type.package.full_path}/property_{snake_caps(attr.type.name)}')

        if attr.element is not None and not isinstance(attr.element, str):
            imports.add(f'{attr.element.package.full_path}/property_{snake_caps(attr.element.name)}')

    return imports


def render_resource(package, resource, output_path):
    imports = _extract_imports(resource.attributes)
    
    rendered = render_resource_template(
        package_name=package.full_path,
        resource_name=resource.name,
        attributes=resource.attributes,
        imports=imports,
    )

    with open(output_path, 'w+') as file_:
        file_.write(rendered)
    

def render_property(property, output_path):
    imports = _extract_imports(property.attributes)
        
    rendered = render_property_template(
        package_name=property.package.full_path,
        property_name=property.name,
        attributes=property.attributes,
        imports=imports,
    )

    with open(output_path, 'w+') as file_:
        file_.write(rendered)
    

def _handle_deferred_types(package, type):
    if isinstance(type, str) and type.startswith('DEFERRED'):
        deferred_prop = type.split('|')[1]
        
        if deferred_prop not in package.properties:
            if package.parent is not None:
                return _handle_deferred_types(package.parent, type)
            else:
                raise Exception('unable to handle deferred type '+deferred_prop)
        else:
            return package.properties[deferred_prop]
    else:
        return type


def _patch_deferred_types(package, attributes):
    for attr in attributes:
        attr.type = _handle_deferred_types(package, attr.type)
        attr.element = _handle_deferred_types(package, attr.element)


def render_package(package, output_path):
    created_directory = f"{output_path}/{package.name}"
    os.makedirs(created_directory, exist_ok=True)
    
    # patch up deferred properties
    for property in package.properties.values():
        _patch_deferred_types(package, property.attributes)

    for resource in package.resources.values():
        _patch_deferred_types(package, resource.attributes)
    
    # render properties
    for property in package.properties.values():
        file_name = f"property_{snake_caps(property.name)}.go"
        file_path = f"{created_directory}/{file_name}"
        
        render_property(property, file_path)
    
    # render resources
    for resource in package.resources.values():
        file_name = f"resource_{snake_caps(resource.name)}.go"
        file_path = f"{created_directory}/{file_name}"
        
        render_resource(package, resource, file_path)
    
    # print(json.dumps(package.as_dict(), indent=2))
    
    # traverse and render subpackages    
    for subpackage in package.subpackages.values():
        render_package(subpackage, created_directory)
