// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigurationTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationTemplateCreate,
		Read:   resourceConfigurationTemplateRead,
		Update: resourceConfigurationTemplateUpdate,
		Delete: resourceConfigurationTemplateDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"environment_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"option_settings": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTemplateConfigurationOptionSetting(),
				Required: false,
			},
			"platform_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"solution_stack_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"source_configuration": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTemplateSourceConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceConfigurationTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}

func resourceConfigurationTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}

func resourceConfigurationTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}

func resourceConfigurationTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}