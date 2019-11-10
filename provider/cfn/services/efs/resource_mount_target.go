// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package efs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMountTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceMountTargetCreate,
		Read:   resourceMountTargetRead,
		Update: resourceMountTargetUpdate,
		Delete: resourceMountTargetDelete,

		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_address": {
				Type: schema.TypeString,
				Required: false,
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
		},
	}
}

func resourceMountTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EFS::MountTarget", data, meta)
}

func resourceMountTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EFS::MountTarget", data, meta)
}

func resourceMountTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EFS::MountTarget", data, meta)
}

func resourceMountTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EFS::MountTarget", data, meta)
}