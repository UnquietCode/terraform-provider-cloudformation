// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceElasticBeanstalkConfigurationTemplateExists,
		Read: resourceElasticBeanstalkConfigurationTemplateRead,
		Create: resourceElasticBeanstalkConfigurationTemplateCreate,
		Update: resourceElasticBeanstalkConfigurationTemplateUpdate,
		Delete: resourceElasticBeanstalkConfigurationTemplateDelete,
		CustomizeDiff: resourceElasticBeanstalkConfigurationTemplateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"environment_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"option_settings": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTemplateConfigurationOptionSetting(),
				Optional: true,
			},
			"platform_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"solution_stack_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_configuration": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTemplateSourceConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElasticBeanstalkConfigurationTemplateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElasticBeanstalkConfigurationTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticBeanstalk::ConfigurationTemplate", ResourceElasticBeanstalkConfigurationTemplate(), data, meta)
}

func resourceElasticBeanstalkConfigurationTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::ConfigurationTemplate", ResourceElasticBeanstalkConfigurationTemplate(), data, meta)
}

func resourceElasticBeanstalkConfigurationTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::ConfigurationTemplate", ResourceElasticBeanstalkConfigurationTemplate(), data, meta)
}

func resourceElasticBeanstalkConfigurationTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::ConfigurationTemplate", data, meta)
}

func resourceElasticBeanstalkConfigurationTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
