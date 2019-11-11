// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html

package efs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEFSMountTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceEFSMountTargetCreate,
		Read:   resourceEFSMountTargetRead,
		Update: resourceEFSMountTargetUpdate,
		Delete: resourceEFSMountTargetDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"file_system_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				Set: schema.HashString,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEFSMountTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EFS::MountTarget", ResourceEFSMountTarget(), data, meta)
}

func resourceEFSMountTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EFS::MountTarget", ResourceEFSMountTarget(), data, meta)
}

func resourceEFSMountTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EFS::MountTarget", ResourceEFSMountTarget(), data, meta)
}

func resourceEFSMountTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EFS::MountTarget", data, meta)
}