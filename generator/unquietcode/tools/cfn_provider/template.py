import re
from string import Template
from datetime import date
from unquietcode.tools.cfn_provider.models import ComplexType

# TODO make these dynamic
PROVIDER_VERSION = '0.0'

DEAD_LINE = '~x~x~x~x~x~x~x~'

RESOURCE_ATTRIBUTE_TEMPLATE = Template(
"""
			"${name}": {
				Type: ${type},
				Elem: ${elem},
				Required: ${required},
				ForceNew: ${force_replace},
				MaxItems: ${max_items},
				Set: ${set_function},
			},
"""[1:])


def _remove_dead_lines(content):
	lines = []

	for line in content.split('\n'):
		if DEAD_LINE in line:
			continue

		lines.append(line)

	content = '\n'.join(lines)
	content = content[:-1] if len(lines) > 1 else content
	return content


def _render_attribute_template(*, package_name, attribute):
	attribute_type = attribute.type.type
	
	if isinstance(attribute_type, ComplexType):
		at_property_prefix = 'property'
		
		if attribute_type.package.full_path == 'cfn/misc':
			at_property_prefix = 'Property'
				
		attribute_type = f'{at_property_prefix}{attribute_type.name}()'

	attribute_elem = attribute.type.element_type
	
	if attribute_elem is not None and not isinstance(attribute_elem, str):
		package_prefix = ""
		
		if attribute_elem.package.name != package_name:
			package_prefix = attribute_elem.package.name + "."
			
		ae_property_prefix = 'property'
		
		if attribute_elem.package.full_path == 'cfn/misc':
			ae_property_prefix = 'Property'
		
		attribute_elem = f'{package_prefix}{ae_property_prefix}{attribute_elem.name}()'
		
	rendered = RESOURCE_ATTRIBUTE_TEMPLATE.substitute(dict(
		name=attribute.go_symbol,
		type=attribute_type,
		elem=attribute_elem or DEAD_LINE,
		required="true" if attribute.required else "false",
		force_replace="true" if attribute.will_replace else DEAD_LINE,
		max_items=attribute.type.max_items or DEAD_LINE,
		set_function=attribute.type.set_function or DEAD_LINE,
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered


HEADER_TEMPLATE = Template(
"""
// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on ${date}, using version ${provider_version} of the cfn
// terraform provider, and version ${schema_version} of the CloudFormation
// resource specification.
"""[1:-1])


def _header_stanza(cfn_version):
	return HEADER_TEMPLATE.substitute(dict(
		date=date.today().strftime("%d-%m-%Y"),
		provider_version=PROVIDER_VERSION,
		schema_version=cfn_version,
	))
	

def _imports_stanza(imports):
	import_lines = [f'	"github.com/unquietcode/terraform-cfn-provider/{i}"' for i in sorted(imports)]
	import_lines = '\n'.join(import_lines)
	
	return import_lines or DEAD_LINE


RESOURCE_TEMPLATE = Template(
"""
${header}

package ${package}

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
${imports}
)

func Resource${name}() *schema.Resource {
	return &schema.Resource{
		Create: resource${name}Create,
		Read:   resource${name}Read,
		Update: resource${name}Update,
		Delete: resource${name}Delete,

		Schema: map[string]*schema.Schema{
${attributes}
		},
	}
}

func resource${name}Create(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("${cfn_type}", data, meta)
}

func resource${name}Read(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("${cfn_type}", data, meta)
}

func resource${name}Update(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("${cfn_type}", data, meta)
}

func resource${name}Delete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("${cfn_type}", data, meta)
}
"""[1:])



def render_resource_template(*, cfn_version, imports, package_name, resource_name, cfn_type, attributes):
	rendered_attributes = [
		_render_attribute_template(package_name=package_name, attribute=attribute)
		for attribute in attributes
	]
	rendered_attributes = '\n'.join(rendered_attributes)
	
	rendered = RESOURCE_TEMPLATE.substitute(dict(
		header=_header_stanza(cfn_version),
		package=package_name,
		name=resource_name,
		cfn_type=cfn_type,
		attributes=rendered_attributes,
		imports=_imports_stanza(imports)
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered
	


PROPERTY_TEMPLATE = Template(
"""
${header}

package ${package}

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
${imports}
)

func ${prefix}${name}() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
${attributes}
		},
	}
}
"""[1:])


def render_property_template(*, cfn_version, package_name, property_name, attributes, imports):
	rendered_attributes = [
		_render_attribute_template(package_name=package_name, attribute=attribute)
		for attribute in attributes
	]
	rendered_attributes = '\n'.join(rendered_attributes)
	
	rendered = PROPERTY_TEMPLATE.substitute(dict(
		header=_header_stanza(cfn_version),
		package=package_name,
		prefix='property' if property_name != 'Tag' else 'Property',
		name=property_name,
		attributes=rendered_attributes,
		imports=_imports_stanza(imports),
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered


def _resources_stanza(resources):
	resource_lines = [f'			"{r.go_symbol}_resource": {r.package.name}.Resource{r.name}(),' for r in resources]
	resource_lines = '\n'.join(resource_lines)
	return resource_lines or DEAD_LINE


def _datasources_stanza(datasources):
	datasource_lines = [f'			"{d.go_symbol}_data_source": {d.package.name}.Datasource{d.name}(),' for d in datasources]
	datasource_lines = '\n'.join(datasource_lines)
	return datasource_lines or DEAD_LINE

	


PROVIDER_TEMPLATE = Template(
"""
${header}

package ${package}

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
${imports}
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
${datasources}
		},
		ResourcesMap: map[string]*schema.Resource{
${resources}
		},
	}
}
"""[1:])


def render_provider_template(*, cfn_version, package_name, imports, datasources, resources):
	rendered = PROVIDER_TEMPLATE.substitute(dict(
		header=_header_stanza(cfn_version),
		package=package_name,
		imports=_imports_stanza(imports),
		resources=_resources_stanza(resources),
		datasources=_datasources_stanza(datasources),
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered

