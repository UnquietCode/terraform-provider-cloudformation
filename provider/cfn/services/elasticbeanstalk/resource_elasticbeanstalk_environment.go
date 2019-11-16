// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticBeanstalkEnvironmentType string = "AWS::ElasticBeanstalk::Environment"

var elasticBeanstalkEnvironmentProperties map[string]string = map[string]string{
	"application_name": "ApplicationName",
	"cname_prefix": "CNAMEPrefix",
	"description": "Description",
	"environment_name": "EnvironmentName",
	"option_settings": "OptionSettings",
	"platform_arn": "PlatformArn",
	"solution_stack_name": "SolutionStackName",
	"tags": "Tags",
	"template_name": "TemplateName",
	"tier": "Tier",
	"version_label": "VersionLabel",
}

func ResourceElasticBeanstalkEnvironment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElasticBeanstalkEnvironmentExists,
		Read: resourceElasticBeanstalkEnvironmentRead,
		Create: resourceElasticBeanstalkEnvironmentCreate,
		Update: resourceElasticBeanstalkEnvironmentUpdate,
		Delete: resourceElasticBeanstalkEnvironmentDelete,
		CustomizeDiff: resourceElasticBeanstalkEnvironmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"cname_prefix": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"environment_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"option_settings": {
				Type: schema.TypeList,
				Elem: propertyEnvironmentOptionSetting(),
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"template_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tier": {
				Type: schema.TypeList,
				Elem: propertyEnvironmentTier(),
				Optional: true,
				MaxItems: 1,
			},
			"version_label": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElasticBeanstalkEnvironmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElasticBeanstalkEnvironmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticBeanstalkEnvironmentType, ResourceElasticBeanstalkEnvironment(), data, meta)
}

func resourceElasticBeanstalkEnvironmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elasticBeanstalkEnvironmentType, ResourceElasticBeanstalkEnvironment(), data, elasticBeanstalkEnvironmentProperties, meta)
}

func resourceElasticBeanstalkEnvironmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elasticBeanstalkEnvironmentType, ResourceElasticBeanstalkEnvironment(), data, elasticBeanstalkEnvironmentProperties, meta)
}

func resourceElasticBeanstalkEnvironmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elasticBeanstalkEnvironmentType, data, meta)
}

func resourceElasticBeanstalkEnvironmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elasticBeanstalkEnvironmentType, data, meta)
}
