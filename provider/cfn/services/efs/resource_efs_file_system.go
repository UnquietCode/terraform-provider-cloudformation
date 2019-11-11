// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html

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
				Optional: true,
				ForceNew: true,
			},
			"file_system_tags": {
				Type: schema.TypeSet,
				Elem: propertyFileSystemElasticFileSystemTag(),
				Optional: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"lifecycle_policies": {
				Type: schema.TypeSet,
				Elem: propertyFileSystemLifecyclePolicy(),
				Optional: true,
			},
			"performance_mode": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"provisioned_throughput_in_mibps": {
				Type: schema.TypeFloat,
				Optional: true,
			},
			"throughput_mode": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEFSFileSystemCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EFS::FileSystem", ResourceEFSFileSystem(), data, meta)
}

func resourceEFSFileSystemRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EFS::FileSystem", ResourceEFSFileSystem(), data, meta)
}

func resourceEFSFileSystemUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EFS::FileSystem", ResourceEFSFileSystem(), data, meta)
}

func resourceEFSFileSystemDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EFS::FileSystem", data, meta)
}