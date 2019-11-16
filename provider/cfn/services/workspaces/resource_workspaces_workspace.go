// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html

package workspaces

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const workSpacesWorkspaceType string = "AWS::WorkSpaces::Workspace"

var workSpacesWorkspaceProperties map[string]string = map[string]string{
	"bundle_id": "BundleId",
	"directory_id": "DirectoryId",
	"root_volume_encryption_enabled": "RootVolumeEncryptionEnabled",
	"tags": "Tags",
	"user_name": "UserName",
	"user_volume_encryption_enabled": "UserVolumeEncryptionEnabled",
	"volume_encryption_key": "VolumeEncryptionKey",
	"workspace_properties": "WorkspaceProperties",
}

func ResourceWorkSpacesWorkspace() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWorkSpacesWorkspaceExists,
		Read: resourceWorkSpacesWorkspaceRead,
		Create: resourceWorkSpacesWorkspaceCreate,
		Update: resourceWorkSpacesWorkspaceUpdate,
		Delete: resourceWorkSpacesWorkspaceDelete,
		CustomizeDiff: resourceWorkSpacesWorkspaceCustomizeDiff,
		
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
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"user_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"user_volume_encryption_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"volume_encryption_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"workspace_properties": {
				Type: schema.TypeSet,
				Elem: propertyWorkspaceWorkspaceProperties(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWorkSpacesWorkspaceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWorkSpacesWorkspaceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(workSpacesWorkspaceType, ResourceWorkSpacesWorkspace(), data, meta)
}

func resourceWorkSpacesWorkspaceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(workSpacesWorkspaceType, ResourceWorkSpacesWorkspace(), data, workSpacesWorkspaceProperties, meta)
}

func resourceWorkSpacesWorkspaceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(workSpacesWorkspaceType, ResourceWorkSpacesWorkspace(), data, workSpacesWorkspaceProperties, meta)
}

func resourceWorkSpacesWorkspaceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(workSpacesWorkspaceType, data, meta)
}

func resourceWorkSpacesWorkspaceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(workSpacesWorkspaceType, data, meta)
}
