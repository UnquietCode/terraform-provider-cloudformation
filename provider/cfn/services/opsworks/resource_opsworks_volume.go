// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-volume.html

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOpsWorksVolume() *schema.Resource {
	return &schema.Resource{
		Exists: resourceOpsWorksVolumeExists,
		Read: resourceOpsWorksVolumeRead,
		Create: resourceOpsWorksVolumeCreate,
		Update: resourceOpsWorksVolumeUpdate,
		Delete: resourceOpsWorksVolumeDelete,
		CustomizeDiff: resourceOpsWorksVolumeCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"ec2_volume_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"mount_point": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"stack_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceOpsWorksVolumeExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceOpsWorksVolumeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::Volume", ResourceOpsWorksVolume(), data, meta)
}

func resourceOpsWorksVolumeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::Volume", ResourceOpsWorksVolume(), data, meta)
}

func resourceOpsWorksVolumeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::Volume", ResourceOpsWorksVolume(), data, meta)
}

func resourceOpsWorksVolumeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::Volume", data, meta)
}

func resourceOpsWorksVolumeCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::OpsWorks::Volume", data, meta)
}

