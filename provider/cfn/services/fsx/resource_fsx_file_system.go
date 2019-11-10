// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html

package fsx

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceFSxFileSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceFSxFileSystemCreate,
		Read:   resourceFSxFileSystemRead,
		Update: resourceFSxFileSystemUpdate,
		Delete: resourceFSxFileSystemDelete,

		Schema: map[string]*schema.Schema{
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"storage_capacity": {
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"file_system_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lustre_configuration": {
				Type: schema.TypeList,
				Elem: propertyFileSystemLustreConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"backup_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"windows_configuration": {
				Type: schema.TypeList,
				Elem: propertyFileSystemWindowsConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceFSxFileSystemCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::FSx::FileSystem", data, meta)
}

func resourceFSxFileSystemRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::FSx::FileSystem", data, meta)
}

func resourceFSxFileSystemUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::FSx::FileSystem", data, meta)
}

func resourceFSxFileSystemDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::FSx::FileSystem", data, meta)
}