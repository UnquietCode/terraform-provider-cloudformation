// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotanalytics-channel.html

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTAnalyticsChannelType string = "AWS::IoTAnalytics::Channel"

func ResourceIoTAnalyticsChannel() *schema.Resource {
	return &schema.Resource{
		Read: resourceIoTAnalyticsChannelRead,
		Create: resourceIoTAnalyticsChannelCreate,
		Update: resourceIoTAnalyticsChannelUpdate,
		Delete: resourceIoTAnalyticsChannelDelete,
		CustomizeDiff: resourceIoTAnalyticsChannelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"channel_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"channel_storage": {
				Type: schema.TypeSet,
				Elem: propertyChannelChannelStorage(),
				Optional: true,
				MaxItems: 1,
			},
			"retention_period": {
				Type: schema.TypeSet,
				Elem: propertyChannelRetentionPeriod(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTAnalyticsChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTAnalyticsChannelType, ResourceIoTAnalyticsChannel(), data, meta)
}

func resourceIoTAnalyticsChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTAnalyticsChannelType, ResourceIoTAnalyticsChannel(), data, meta)
}

func resourceIoTAnalyticsChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTAnalyticsChannelType, ResourceIoTAnalyticsChannel(), data, meta)
}

func resourceIoTAnalyticsChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTAnalyticsChannelType, data, meta)
}

func resourceIoTAnalyticsChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTAnalyticsChannelType, data, meta)
}
