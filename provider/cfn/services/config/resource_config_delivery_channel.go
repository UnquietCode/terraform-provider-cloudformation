// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const configDeliveryChannelType string = "AWS::Config::DeliveryChannel"

var configDeliveryChannelProperties map[string]string = map[string]string{
	"config_snapshot_delivery_properties": "ConfigSnapshotDeliveryProperties",
	"name": "Name",
	"s3_bucket_name": "S3BucketName",
	"s3_key_prefix": "S3KeyPrefix",
	"sns_topic_arn": "SnsTopicARN",
}

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
				Type: schema.TypeSet,
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
	return plugin.ResourceRead(configDeliveryChannelType, ResourceConfigDeliveryChannel(), data, meta)
}

func resourceConfigDeliveryChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(configDeliveryChannelType, ResourceConfigDeliveryChannel(), data, configDeliveryChannelProperties, meta)
}

func resourceConfigDeliveryChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(configDeliveryChannelType, ResourceConfigDeliveryChannel(), data, configDeliveryChannelProperties, meta)
}

func resourceConfigDeliveryChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(configDeliveryChannelType, data, meta)
}

func resourceConfigDeliveryChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(configDeliveryChannelType, data, meta)
}
