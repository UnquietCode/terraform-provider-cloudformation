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

func ResourceApplicationVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationVersionCreate,
		Read:   resourceApplicationVersionRead,
		Update: resourceApplicationVersionUpdate,
		Delete: resourceApplicationVersionDelete,

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

func resourceApplicationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticBeanstalk::ApplicationVersion", data, meta)
}

func resourceApplicationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticBeanstalk::ApplicationVersion", data, meta)
}

func resourceApplicationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticBeanstalk::ApplicationVersion", data, meta)
}

func resourceApplicationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticBeanstalk::ApplicationVersion", data, meta)
}