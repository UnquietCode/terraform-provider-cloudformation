// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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

func ResourceElasticBeanstalkEnvironment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElasticBeanstalkEnvironmentExists,
		Read:   resourceElasticBeanstalkEnvironmentRead,
		Create: resourceElasticBeanstalkEnvironmentCreate,
		Update: resourceElasticBeanstalkEnvironmentUpdate,
		Delete: resourceElasticBeanstalkEnvironmentDelete,
		
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
	return plugin.ResourceRead("AWS::ElasticBeanstalk::Environment", ResourceElasticBeanstalkEnvironment(), data, meta)
}

func resourceElasticBeanstalkEnvironmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::Environment", ResourceElasticBeanstalkEnvironment(), data, meta)
}

func resourceElasticBeanstalkEnvironmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::Environment", ResourceElasticBeanstalkEnvironment(), data, meta)
}

func resourceElasticBeanstalkEnvironmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::Environment", data, meta)
}