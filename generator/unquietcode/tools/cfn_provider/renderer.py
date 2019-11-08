import os
import json
from itertools import chain

from .template import render_resource_template, render_property_template, render_provider_template
from unquietcode.tools.cfn_provider.models import ComplexType
from unquietcode.tools.cfn_provider.utils import snake_caps


def _extract_imports(attributes):
    imports = set()
    
    for attr in attributes:
        if isinstance(attr.type.type, ComplexType):
            imports.add(attr.type.type.package.full_path)

        if isinstance(attr.type.element_type, ComplexType):
            imports.add(attr.type.element_type.package.full_path)

    return imports


def render_provider(provider_data, output_path):
    package = provider_data.top_level_package
    
    # do the normal rendering first
    render_package(provider_data.cfn_version, package, output_path)
    
    # gather up all the imports, resources, etc
    imports = set()
    packages = [package]
    resources = []
    datasources = []
    properties = []
    
    while len(packages) > 0:
        p = packages.pop()
        
        if len(p.resources) > 0 or len(p.datasources) > 0:
            imports.add(p.full_path)
        
        resources.extend(p.resources.values())
        datasources.extend(p.datasources.values())
        
        packages.extend(p.subpackages.values())
    
    resources = sorted(resources, key=lambda _: (_.package.full_path, _.name))
    datasources = sorted(datasources, key=lambda _: (_.package.full_path, _.name))
    properties = sorted(properties, key=lambda _: (_.package.full_path, _.name))
    
    # we don't need the import for our own package
    imports.discard(package.full_path)
    
    rendered = render_provider_template(
        cfn_version=provider_data.cfn_version,
        package_name=package.name,
        resources=resources,
        datasources=datasources,
        imports=imports,
    )

    with open(f"{output_path}/{package.full_path}/provider.go", 'w+') as file_:
        file_.write(rendered)


def render_resource(cfn_version, resource, output_path):
    imports = _extract_imports(resource.attributes.values())
    # imports.add('cfn')  # ensure top level package is imported
    imports.add('plugin')
    imports.discard(resource.package.full_path)  # remove self
    
    rendered = render_resource_template(
        cfn_version=cfn_version,
        package_name=resource.package.name,
        resource_name=resource.name,
        cfn_type=resource.cfn_type,
        attributes=resource.attributes.values(),
        imports=imports,
    )

    with open(output_path, 'w+') as file_:
        file_.write(rendered)
    

def render_property(cfn_version, property, output_path):
    imports = _extract_imports(property.attributes.values())
    imports.discard(property.package.full_path)  # remove self
    
    rendered = render_property_template(
        cfn_version=cfn_version,
        package_name=property.package.name,
        property_name=property.name,
        attributes=property.attributes.values(),
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
            
            # TODO total hack
            elif 'misc' in package.subpackages:
                misc_package = package.subpackages['misc']
                
                if deferred_prop in misc_package.properties:
                    return misc_package.properties[deferred_prop]
                else:
                    raise Exception('unable to handle deferred type '+deferred_prop)
            else:
                raise Exception('unable to handle deferred type '+deferred_prop)
        else:
            return package.properties[deferred_prop]
    else:
        return type


def _patch_deferred_types(package, attributes):
    for attr in attributes.values():
        attr.type.type = _handle_deferred_types(package, attr.type.type)
        attr.type.element_type = _handle_deferred_types(package, attr.type.element_type)


def render_package(cfn_version, package, output_path):
    created_directory = f"{output_path}/{package.name}"
    os.makedirs(created_directory, exist_ok=True)
    
    # patch up deferred properties
    for property in package.properties.values():
        _patch_deferred_types(package, property.attributes)

    for resource in package.resources.values():
        _patch_deferred_types(package, resource.attributes)
    
    # render properties
    for property in package.properties.values():
        file_name = f"property_{property.go_symbol}.go"
        file_path = f"{created_directory}/{file_name}"
        render_property(cfn_version, property, file_path)
    
    # render resources
    for resource in package.resources.values():
        file_name = f"resource_{resource.go_symbol}.go"
        file_path = f"{created_directory}/{file_name}"
        render_resource(cfn_version, resource, file_path)
    
    # print(json.dumps(package.as_dict(), indent=2))
    
    # traverse and render subpackages
    sorted_subpackages = sorted(package.subpackages.values(), key=lambda _: _.name)
    
    for subpackage in sorted_subpackages:
        render_package(cfn_version, subpackage, created_directory)
