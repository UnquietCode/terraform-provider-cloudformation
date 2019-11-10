// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApplicationOutput() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationOutputCreate,
		Read:   resourceApplicationOutputRead,
		Update: resourceApplicationOutputUpdate,
		Delete: resourceApplicationOutputDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"output": {
				Type: schema.TypeList,
				Elem: propertyApplicationOutputOutput(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceApplicationOutputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalyticsV2::ApplicationOutput", data, meta)
}

func resourceApplicationOutputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalyticsV2::ApplicationOutput", data, meta)
}

func resourceApplicationOutputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalyticsV2::ApplicationOutput", data, meta)
}

func resourceApplicationOutputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalyticsV2::ApplicationOutput", data, meta)
}