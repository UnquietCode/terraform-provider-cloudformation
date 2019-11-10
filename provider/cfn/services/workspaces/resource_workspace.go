// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package workspaces

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWorkspace() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkspaceCreate,
		Read:   resourceWorkspaceRead,
		Update: resourceWorkspaceUpdate,
		Delete: resourceWorkspaceDelete,

		Schema: map[string]*schema.Schema{
			"bundle_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"directory_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"root_volume_encryption_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_volume_encryption_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"volume_encryption_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"workspace_properties": {
				Type: schema.TypeList,
				Elem: propertyWorkspaceWorkspaceProperties(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceWorkspaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WorkSpaces::Workspace", data, meta)
}

func resourceWorkspaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WorkSpaces::Workspace", data, meta)
}

func resourceWorkspaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WorkSpaces::Workspace", data, meta)
}

func resourceWorkspaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WorkSpaces::Workspace", data, meta)
}