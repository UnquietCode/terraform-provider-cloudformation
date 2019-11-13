// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volume.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2Volume() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VolumeExists,
		Read:   resourceEC2VolumeRead,
		Create: resourceEC2VolumeCreate,
		Update: resourceEC2VolumeUpdate,
		Delete: resourceEC2VolumeDelete,
		CustomizeDiff: resourceEC2VolumeCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"auto_enable_io": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: true,
			},
			"encrypted": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"iops": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"snapshot_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"volume_type": {
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

func resourceEC2VolumeExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2VolumeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::Volume", ResourceEC2Volume(), data, meta)
}

func resourceEC2VolumeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::Volume", ResourceEC2Volume(), data, meta)
}

func resourceEC2VolumeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::Volume", ResourceEC2Volume(), data, meta)
}

func resourceEC2VolumeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::Volume", data, meta)
}

func resourceEC2VolumeCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}