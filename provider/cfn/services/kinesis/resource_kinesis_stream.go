// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html

package kinesis

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisStreamType string = "AWS::Kinesis::Stream"

var kinesisStreamProperties map[string]string = map[string]string{
	"name": "Name",
	"retention_period_hours": "RetentionPeriodHours",
	"shard_count": "ShardCount",
	"stream_encryption": "StreamEncryption",
	"tags": "Tags",
}

func ResourceKinesisStream() *schema.Resource {
	return &schema.Resource{
		Exists: resourceKinesisStreamExists,
		Read: resourceKinesisStreamRead,
		Create: resourceKinesisStreamCreate,
		Update: resourceKinesisStreamUpdate,
		Delete: resourceKinesisStreamDelete,
		CustomizeDiff: resourceKinesisStreamCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"retention_period_hours": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"shard_count": {
				Type: schema.TypeInt,
				Required: true,
			},
			"stream_encryption": {
				Type: schema.TypeSet,
				Elem: propertyStreamStreamEncryption(),
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

func resourceKinesisStreamExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKinesisStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisStreamType, ResourceKinesisStream(), data, meta)
}

func resourceKinesisStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kinesisStreamType, ResourceKinesisStream(), data, kinesisStreamProperties, meta)
}

func resourceKinesisStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kinesisStreamType, ResourceKinesisStream(), data, kinesisStreamProperties, meta)
}

func resourceKinesisStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kinesisStreamType, data, meta)
}

func resourceKinesisStreamCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kinesisStreamType, data, meta)
}
