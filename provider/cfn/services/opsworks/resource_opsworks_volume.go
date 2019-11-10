// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceOpsWorksVolumeCreate,
		Read:   resourceOpsWorksVolumeRead,
		Update: resourceOpsWorksVolumeUpdate,
		Delete: resourceOpsWorksVolumeDelete,

		Schema: map[string]*schema.Schema{
			"ec2_volume_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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
				ForceNew: true,
			},
		},
	}
}

func resourceOpsWorksVolumeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::OpsWorks::Volume", data, meta)
}

func resourceOpsWorksVolumeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::OpsWorks::Volume", data, meta)
}

func resourceOpsWorksVolumeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::OpsWorks::Volume", data, meta)
}

func resourceOpsWorksVolumeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::OpsWorks::Volume", data, meta)
}