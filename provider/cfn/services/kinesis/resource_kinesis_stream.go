// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
				Type: schema.TypeList,
				Elem: propertyStreamStreamEncryption(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceKinesisStreamExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKinesisStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Kinesis::Stream", ResourceKinesisStream(), data, meta)
}

func resourceKinesisStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Kinesis::Stream", ResourceKinesisStream(), data, meta)
}

func resourceKinesisStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Kinesis::Stream", ResourceKinesisStream(), data, meta)
}

func resourceKinesisStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Kinesis::Stream", data, meta)
}

func resourceKinesisStreamCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
