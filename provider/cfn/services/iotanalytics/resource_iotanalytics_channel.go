// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTAnalyticsChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTAnalyticsChannelCreate,
		Read:   resourceIoTAnalyticsChannelRead,
		Update: resourceIoTAnalyticsChannelUpdate,
		Delete: resourceIoTAnalyticsChannelDelete,

		Schema: map[string]*schema.Schema{
			"channel_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"channel_storage": {
				Type: schema.TypeList,
				Elem: propertyChannelChannelStorage(),
				Required: false,
				MaxItems: 1,
			},
			"retention_period": {
				Type: schema.TypeList,
				Elem: propertyChannelRetentionPeriod(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceIoTAnalyticsChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoTAnalytics::Channel", data, meta)
}

func resourceIoTAnalyticsChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoTAnalytics::Channel", data, meta)
}

func resourceIoTAnalyticsChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoTAnalytics::Channel", data, meta)
}

func resourceIoTAnalyticsChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoTAnalytics::Channel", data, meta)
}