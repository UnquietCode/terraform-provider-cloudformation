// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCreate,
		Read:   resourceApplicationRead,
		Update: resourceApplicationUpdate,
		Delete: resourceApplicationDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"inputs": {
				Type: schema.TypeList,
				Elem: propertyApplicationInput(),
				Required: true,
			},
			"application_description": {
				Type: schema.TypeString,
				Required: false,
			},
			"application_code": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceApplicationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalytics::Application", data, meta)
}

func resourceApplicationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalytics::Application", data, meta)
}

func resourceApplicationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalytics::Application", data, meta)
}

func resourceApplicationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalytics::Application", data, meta)
}