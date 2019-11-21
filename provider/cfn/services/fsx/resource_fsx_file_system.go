// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html

package fsx

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const fSxFileSystemType string = "AWS::FSx::FileSystem"

func ResourceFSxFileSystem() *schema.Resource {
	return &schema.Resource{
		Read: resourceFSxFileSystemRead,
		Create: resourceFSxFileSystemCreate,
		Update: resourceFSxFileSystemUpdate,
		Delete: resourceFSxFileSystemDelete,
		CustomizeDiff: resourceFSxFileSystemCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"storage_capacity": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"file_system_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"lustre_configuration": {
				Type: schema.TypeSet,
				Elem: propertyFileSystemLustreConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"backup_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"windows_configuration": {
				Type: schema.TypeSet,
				Elem: propertyFileSystemWindowsConfiguration(),
				Optional: true,
				MaxItems: 1,
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

func resourceFSxFileSystemRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(fSxFileSystemType, ResourceFSxFileSystem(), data, meta)
}

func resourceFSxFileSystemCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(fSxFileSystemType, ResourceFSxFileSystem(), data, meta)
}

func resourceFSxFileSystemUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(fSxFileSystemType, ResourceFSxFileSystem(), data, meta)
}

func resourceFSxFileSystemDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(fSxFileSystemType, data, meta)
}

func resourceFSxFileSystemCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(fSxFileSystemType, data, meta)
}
