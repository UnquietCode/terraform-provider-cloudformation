// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html

package config

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const configDeliveryChannelType string = "AWS::Config::DeliveryChannel"

func DatasourceConfigDeliveryChannel() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConfigDeliveryChannelRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceConfigDeliveryChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(configDeliveryChannelType, DatasourceConfigDeliveryChannel(), data, meta)
}
