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


RESOURCE_TEMPLATE = Template(
"""
package ${package}

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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


def _render_attribute_template(*, attribute):
	rendered = RESOURCE_ATTRIBUTE_TEMPLATE.substitute(dict(
		name=attribute.name,
		type=attribute.type,
		elem=attribute.element or DEAD_LINE,
		required="true" if attribute.required else "false",
		force_replace="true" if attribute.will_replace else DEAD_LINE,
	))

	lines = []

	for line in rendered.splitlines():
		if DEAD_LINE in line:
			continue

		lines.append(line)

	rendered = '\n'.join(lines)
	return rendered


def render_resource_template(*, package_name, resource_name, attributes):
    rendered_attributes = [
        _render_attribute_template(attribute=attribute)
        for attribute in attributes
    ]
    rendered_attributes = '\n'.join(rendered_attributes)

    return RESOURCE_TEMPLATE.substitute(dict(
        package=package_name,
        name=resource_name,
        attributes=rendered_attributes,
    ))
