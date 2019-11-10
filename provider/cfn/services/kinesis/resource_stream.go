// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesis

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStream() *schema.Resource {
	return &schema.Resource{
		Create: resourceStreamCreate,
		Read:   resourceStreamRead,
		Update: resourceStreamUpdate,
		Delete: resourceStreamDelete,

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
				Elem: propertyStreamStreamEncryption(),
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

func resourceStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Kinesis::Stream", data, meta)
}

func resourceStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Kinesis::Stream", data, meta)
}

func resourceStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Kinesis::Stream", data, meta)
}

func resourceStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Kinesis::Stream", data, meta)
}