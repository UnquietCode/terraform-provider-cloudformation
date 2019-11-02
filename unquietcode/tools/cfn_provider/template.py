import re
from string import Template

DEAD_LINE = '~x~x~x~x~x~x~x~'

RESOURCE_ATTRIBUTE_TEMPLATE = Template(
"""
			"${name}": {
				Type: ${type},
				Elem: ${elem},
				Required: ${required},
				ForceNew: ${force_replace},
			},
"""[1:])


def _remove_dead_lines(content):
	lines = []

	for line in content.split('\n'):
		if DEAD_LINE in line:
			continue

		lines.append(line)

	content = '\n'.join(lines)
	return content


def _render_attribute_template(*, attribute):
	rendered = RESOURCE_ATTRIBUTE_TEMPLATE.substitute(dict(
		name=attribute.name,
		type=attribute.type,
		elem=attribute.element or DEAD_LINE,
		required="true" if attribute.required else "false",
		force_replace="true" if attribute.will_replace else DEAD_LINE,
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered



RESOURCE_TEMPLATE = Template(
"""
package ${package}

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
${imports}
)

func resource${name}() *schema.Resource {
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

func resource${name}Create(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource${name}Read(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource${name}Update(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resource${name}Delete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
"""[1:])



def render_resource_template(*, imports, package_name, resource_name, attributes):
	rendered_attributes = [
		_render_attribute_template(attribute=attribute)
		for attribute in attributes
	]
	rendered_attributes = '\n'.join(rendered_attributes)
	
	import_lines = [f'	"github.com/unquietcode/cfn-provider/{package_name}/{i}"' for i in imports]
	import_lines = '\n'.join(import_lines)
	
	rendered = RESOURCE_TEMPLATE.substitute(dict(
		package=package_name,
		name=resource_name,
		attributes=rendered_attributes,
		imports=import_lines or DEAD_LINE,
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered
	


PROPERTY_TEMPLATE = Template(
"""
package ${package}

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func property${name}() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
${attributes}
		},
	}
}
"""[1:])


def render_property_template(*, package_name, property_name, attributes):
	rendered_attributes = [
		_render_attribute_template(attribute=attribute)
		for attribute in attributes
	]
	rendered_attributes = '\n'.join(rendered_attributes)
	
	rendered = PROPERTY_TEMPLATE.substitute(dict(
		package=package_name,
		name=property_name,
		attributes=rendered_attributes,
	))
	
	return rendered
