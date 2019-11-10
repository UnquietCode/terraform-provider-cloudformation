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

func ResourceEFSFileSystem() *schema.Resource {
	return &schema.Resource{
		Create: resourceEFSFileSystemCreate,
		Read:   resourceEFSFileSystemRead,
		Update: resourceEFSFileSystemUpdate,
		Delete: resourceEFSFileSystemDelete,

		Schema: map[string]*schema.Schema{
			"encrypted": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"file_system_tags": {
				Type: schema.TypeSet,
				Elem: propertyFileSystemElasticFileSystemTag(),
				Required: false,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"lifecycle_policies": {
				Type: schema.TypeSet,
				Elem: propertyFileSystemLifecyclePolicy(),
				Required: false,
			},
			"performance_mode": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"provisioned_throughput_in_mibps": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"throughput_mode": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceEFSFileSystemCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EFS::FileSystem", data, meta)
}

func resourceEFSFileSystemRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EFS::FileSystem", data, meta)
}

func resourceEFSFileSystemUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EFS::FileSystem", data, meta)
}

func resourceEFSFileSystemDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EFS::FileSystem", data, meta)
}