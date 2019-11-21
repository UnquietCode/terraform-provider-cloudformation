// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html

package efs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eFSMountTargetType string = "AWS::EFS::MountTarget"

func ResourceEFSMountTarget() *schema.Resource {
	return &schema.Resource{
		Read: resourceEFSMountTargetRead,
		Create: resourceEFSMountTargetCreate,
		Update: resourceEFSMountTargetUpdate,
		Delete: resourceEFSMountTargetDelete,
		CustomizeDiff: resourceEFSMountTargetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"ip_address": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEFSMountTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eFSMountTargetType, ResourceEFSMountTarget(), data, meta)
}

func resourceEFSMountTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eFSMountTargetType, ResourceEFSMountTarget(), data, meta)
}

func resourceEFSMountTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eFSMountTargetType, ResourceEFSMountTarget(), data, meta)
}

func resourceEFSMountTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eFSMountTargetType, data, meta)
}

func resourceEFSMountTargetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eFSMountTargetType, data, meta)
}
