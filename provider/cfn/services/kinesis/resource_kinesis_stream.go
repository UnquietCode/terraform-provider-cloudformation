// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesis

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisStream() *schema.Resource {
	return &schema.Resource{
		Create: resourceKinesisStreamCreate,
		Read:   resourceKinesisStreamRead,
		Update: resourceKinesisStreamUpdate,
		Delete: resourceKinesisStreamDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"retention_period_hours": {
				Type: schema.TypeInt,
				Required: false,
			},
			"shard_count": {
				Type: schema.TypeInt,
				Required: true,
			},
			"stream_encryption": {
				Type: schema.TypeList,
				Elem: propertyStreamEncryption(),
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

func resourceKinesisStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Kinesis::Stream", data, meta)
}

func resourceKinesisStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Kinesis::Stream", data, meta)
}

func resourceKinesisStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Kinesis::Stream", data, meta)
}

func resourceKinesisStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Kinesis::Stream", data, meta)
}