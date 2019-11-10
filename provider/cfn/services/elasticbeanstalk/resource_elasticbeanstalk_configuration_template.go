// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticBeanstalkConfigurationTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticBeanstalkConfigurationTemplateCreate,
		Read:   resourceElasticBeanstalkConfigurationTemplateRead,
		Update: resourceElasticBeanstalkConfigurationTemplateUpdate,
		Delete: resourceElasticBeanstalkConfigurationTemplateDelete,

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

func resourceElasticBeanstalkConfigurationTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}

func resourceElasticBeanstalkConfigurationTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}

func resourceElasticBeanstalkConfigurationTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}

func resourceElasticBeanstalkConfigurationTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}