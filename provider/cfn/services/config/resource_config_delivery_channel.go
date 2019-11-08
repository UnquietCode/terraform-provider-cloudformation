// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigDeliveryChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigDeliveryChannelCreate,
		Read:   resourceConfigDeliveryChannelRead,
		Update: resourceConfigDeliveryChannelUpdate,
		Delete: resourceConfigDeliveryChannelDelete,

		Schema: map[string]*schema.Schema{
			"config_snapshot_delivery_properties": {
				Type: schema.TypeList,
				Elem: propertyConfigSnapshotDeliveryProperties(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"s3_bucket_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_key_prefix": {
				Type: schema.TypeString,
				Required: false,
			},
			"sns_topic_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceConfigDeliveryChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::DeliveryChannel", data, meta)
}

func resourceConfigDeliveryChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::DeliveryChannel", data, meta)
}

func resourceConfigDeliveryChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::DeliveryChannel", data, meta)
}

func resourceConfigDeliveryChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::DeliveryChannel", data, meta)
}