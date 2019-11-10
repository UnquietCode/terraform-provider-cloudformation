import re
from string import Template
from datetime import date
from unquietcode.tools.cfn_provider.models import ComplexType
from unquietcode.tools.cfn_provider.utils import snake_caps

# TODO make this dynamic
PROVIDER_VERSION = '0.0'

DEAD_LINE = '~x~x~x~x~x~x~x~'
MAX_PROPERTY_RECURSION = 5
PROJECT_URL = "https://github.com/UnquietCode/terraform-provider-cloudformation"


RESOURCE_ATTRIBUTE_TEMPLATE = Template(
"""
			"${name}": {
				Type: ${type},
				Elem: ${elem},
				Required: ${required},
				Optional: ${optional},
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


def _render_attribute_template(*, package_name, schema_name, attribute):
	attribute_type = attribute.type.type
	
	if isinstance(attribute_type, ComplexType):
		at_property_prefix = 'property'
		
		if attribute_type.package.full_path == 'cfn/misc':
			at_property_prefix = 'Property'
		
		if attribute_type == f"property{schema_name}":
			print("recursion")
		
		attribute_type = f'{at_property_prefix}{attribute_type.name}()'

	attribute_elem = attribute.type.element_type
	
	if attribute_elem is not None and not isinstance(attribute_elem, str):
		package_prefix = ""
		recursive = False
		
		if attribute_elem.package.name != package_name:
			package_prefix = attribute_elem.package.name + "."
		else:
			if attribute_elem.name == schema_name:
				recursive = True
	
		ae_property_prefix = 'property'
		
		if attribute_elem.package.full_path == 'cfn/misc':
			ae_property_prefix = 'Property'
		
		call = '()' if recursive is not True else '(strconv.Itoa(int(count + 1)))'
		attribute_elem = f'{package_prefix}{ae_property_prefix}{attribute_elem.name}{call}'
	
	rendered = RESOURCE_ATTRIBUTE_TEMPLATE.substitute(dict(
		name=snake_caps(attribute.name),
		type=attribute_type,
		elem=attribute_elem or DEAD_LINE,
		required="true" if attribute.required is True else DEAD_LINE,
		optional="true" if attribute.required is not True else DEAD_LINE,
		force_replace="true" if attribute.will_replace else DEAD_LINE,
		max_items=attribute.type.max_items or DEAD_LINE,
		set_function=attribute.type.set_function or DEAD_LINE,
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered


HEADER_TEMPLATE = Template(
"""
// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on ${date}, using version ${provider_version} of the cfn terraform provider,
// and version ${schema_version} of the CloudFormation resource specification.
//
// For more information, visit:
//   ${documentation_link}
"""[1:-1])


def _header_stanza(cfn_version, documentation_link):
	return HEADER_TEMPLATE.substitute(dict(
		date=date.today().strftime("%d-%m-%Y"),
		provider_version=PROVIDER_VERSION,
		schema_version=cfn_version,
		documentation_link=documentation_link or PROJECT_URL,
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
		Update: ${update_line},
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



def render_resource_template(*, cfn_version, imports, package_name, resource_name, cfn_type, attributes, documentation_link):
	rendered_attributes = [
		_render_attribute_template(package_name=package_name, schema_name=None, attribute=attribute)
		for attribute in attributes
	]
	rendered_attributes = '\n'.join(rendered_attributes)
	
	# Terraform requires that we check for the case of non-updatable resources
	# TODO computable
	updatable = False
	
	# at least one attribute will not force a replacement
	for attribute in attributes:
		if attribute.will_replace is not True:
			updatable = True
			break
			
	rendered = RESOURCE_TEMPLATE.substitute(dict(
		header=_header_stanza(cfn_version, documentation_link),
		package=package_name,
		name=resource_name,
		cfn_type=cfn_type,
		attributes=rendered_attributes,
		update_line=f"resource{resource_name}Update" if updatable else DEAD_LINE,
		imports=_imports_stanza(imports)
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered
	


PROPERTY_TEMPLATE = Template(
"""
${header}

package ${package}

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
${imports}
)

func ${prefix}${name}(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= ${max_recursion} {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
${attributes}
		},
	}
}
"""[1:])


def render_property_template(*, cfn_version, package_name, property_name, attributes, imports, documentation_link):
	rendered_attributes = [
		_render_attribute_template(package_name=package_name, schema_name=property_name, attribute=attribute)
		for attribute in attributes
	]
	rendered_attributes = '\n'.join(rendered_attributes)
	
	rendered = PROPERTY_TEMPLATE.substitute(dict(
		header=_header_stanza(cfn_version, documentation_link),
		package=package_name,
		prefix='property' if property_name != 'Tag' else 'Property',
		name=property_name,
		attributes=rendered_attributes,
		imports=_imports_stanza(imports),
		max_recursion=MAX_PROPERTY_RECURSION,
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered


def _resources_stanza(resources):
	resource_lines = [f'			"cfn_{snake_caps(r.name)}": {r.package.name}.Resource{r.name}(),' for r in resources]
	resource_lines = '\n'.join(resource_lines)
	return resource_lines or DEAD_LINE


def _datasources_stanza(datasources):
	datasource_lines = [f'			"cfn_{snake_caps(d.name)}": {d.package.name}.Datasource{d.name}(),' for d in datasources]
	datasource_lines = '\n'.join(datasource_lines)
	return datasource_lines or DEAD_LINE

	


PROVIDER_TEMPLATE = Template(
"""
${header}

package ${package}

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
${imports}
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: plugin.ProviderConfigure,
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "name of an S3 bucket where we can store template files",
			},
			"bucket_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "path in the bucket under which CFN can store data",
			},
			"stack_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "exact name to use for the CloudFormation stack",
			},
			"role_arn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ARN of a role to be used for interacting with CloudFormation",
			},
		},
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
		header=_header_stanza(cfn_version, None),
		package=package_name,
		imports=_imports_stanza(imports),
		resources=_resources_stanza(resources),
		datasources=_datasources_stanza(datasources),
	))
	
	rendered = _remove_dead_lines(rendered)
	return rendered

