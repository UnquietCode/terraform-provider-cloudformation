// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticBeanstalkEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticBeanstalkEnvironmentCreate,
		Read:   resourceElasticBeanstalkEnvironmentRead,
		Update: resourceElasticBeanstalkEnvironmentUpdate,
		Delete: resourceElasticBeanstalkEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cname_prefix": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"environment_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"option_settings": {
				Type: schema.TypeList,
				Elem: propertyEnvironmentOptionSetting(),
				Required: false,
			},
			"platform_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"solution_stack_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"template_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"tier": {
				Type: schema.TypeList,
				Elem: propertyEnvironmentTier(),
				Required: false,
				MaxItems: 1,
			},
			"version_label": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceElasticBeanstalkEnvironmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::Environment", data, meta)
}

func resourceElasticBeanstalkEnvironmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticBeanstalk::Environment", data, meta)
}

func resourceElasticBeanstalkEnvironmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::Environment", data, meta)
}

func resourceElasticBeanstalkEnvironmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::Environment", data, meta)
}