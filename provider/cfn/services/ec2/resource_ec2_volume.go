// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const eC2VolumeType string = "AWS::EC2::Volume"

var eC2VolumeProperties map[string]string = map[string]string{
	"auto_enable_io": "AutoEnableIO",
	"availability_zone": "AvailabilityZone",
	"encrypted": "Encrypted",
	"iops": "Iops",
	"kms_key_id": "KmsKeyId",
	"size": "Size",
	"snapshot_id": "SnapshotId",
	"tags": "Tags",
	"volume_type": "VolumeType",
}

func ResourceEC2Volume() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2VolumeExists,
		Read: resourceEC2VolumeRead,
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
			"tags": misc.PropertyTags(),
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
	return plugin.ResourceRead(eC2VolumeType, ResourceEC2Volume(), data, meta)
}

func resourceEC2VolumeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VolumeType, ResourceEC2Volume(), data, eC2VolumeProperties, meta)
}

func resourceEC2VolumeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VolumeType, ResourceEC2Volume(), data, eC2VolumeProperties, meta)
}

func resourceEC2VolumeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VolumeType, data, meta)
}

func resourceEC2VolumeCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VolumeType, data, meta)
}
