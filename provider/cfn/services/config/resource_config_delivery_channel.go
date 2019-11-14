// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigDeliveryChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourceConfigDeliveryChannelExists,
		Read: resourceConfigDeliveryChannelRead,
		Create: resourceConfigDeliveryChannelCreate,
		Update: resourceConfigDeliveryChannelUpdate,
		Delete: resourceConfigDeliveryChannelDelete,
		CustomizeDiff: resourceConfigDeliveryChannelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"config_snapshot_delivery_properties": {
				Type: schema.TypeList,
				Elem: propertyDeliveryChannelConfigSnapshotDeliveryProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"s3_bucket_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_key_prefix": {
				Type: schema.TypeString,
				Optional: true,
			},
			"sns_topic_arn": {
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

func resourceConfigDeliveryChannelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceConfigDeliveryChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::DeliveryChannel", ResourceConfigDeliveryChannel(), data, meta)
}

func resourceConfigDeliveryChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::DeliveryChannel", ResourceConfigDeliveryChannel(), data, meta)
}

func resourceConfigDeliveryChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::DeliveryChannel", ResourceConfigDeliveryChannel(), data, meta)
}

func resourceConfigDeliveryChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::DeliveryChannel", data, meta)
}

func resourceConfigDeliveryChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

