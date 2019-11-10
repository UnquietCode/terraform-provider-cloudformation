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

func ResourceElasticBeanstalkApplicationVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticBeanstalkApplicationVersionCreate,
		Read:   resourceElasticBeanstalkApplicationVersionRead,
		Update: resourceElasticBeanstalkApplicationVersionUpdate,
		Delete: resourceElasticBeanstalkApplicationVersionDelete,

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
			"source_bundle": {
				Type: schema.TypeList,
				Elem: propertyApplicationVersionSourceBundle(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceElasticBeanstalkApplicationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::ApplicationVersion", data, meta)
}

func resourceElasticBeanstalkApplicationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticBeanstalk::ApplicationVersion", data, meta)
}

func resourceElasticBeanstalkApplicationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::ApplicationVersion", data, meta)
}

func resourceElasticBeanstalkApplicationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::ApplicationVersion", data, meta)
}