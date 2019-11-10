// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceVolume() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolumeCreate,
		Read:   resourceVolumeRead,
		Update: resourceVolumeUpdate,
		Delete: resourceVolumeDelete,

		Schema: map[string]*schema.Schema{
			"auto_enable_io": {
				Type: schema.TypeBool,
				Required: false,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: true,
			},
			"encrypted": {
				Type: schema.TypeBool,
				Required: false,
			},
			"iops": {
				Type: schema.TypeInt,
				Required: false,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"size": {
				Type: schema.TypeInt,
				Required: false,
			},
			"snapshot_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"volume_type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceVolumeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::Volume", data, meta)
}

func resourceVolumeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::Volume", data, meta)
}

func resourceVolumeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::Volume", data, meta)
}

func resourceVolumeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::Volume", data, meta)
}