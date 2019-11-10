// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package fsx

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceFileSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileSystemCreate,
		Read:   resourceFileSystemRead,
		Update: resourceFileSystemUpdate,
		Delete: resourceFileSystemDelete,

		Schema: map[string]*schema.Schema{
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"storage_capacity": {
				Type: schema.TypeInt,
				Required: false,
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
				Required: false,
				MaxItems: 1,
			},
			"backup_id": {
				Type: schema.TypeString,
				Required: false,
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
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"windows_configuration": {
				Type: schema.TypeList,
				Elem: propertyFileSystemWindowsConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceFileSystemCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::FSx::FileSystem", data, meta)
}

func resourceFileSystemRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::FSx::FileSystem", data, meta)
}

func resourceFileSystemUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::FSx::FileSystem", data, meta)
}

func resourceFileSystemDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::FSx::FileSystem", data, meta)
}