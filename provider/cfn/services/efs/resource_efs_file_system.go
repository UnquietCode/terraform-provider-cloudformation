// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html

package efs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eFSFileSystemType string = "AWS::EFS::FileSystem"

func ResourceEFSFileSystem() *schema.Resource {
	return &schema.Resource{
		Read: resourceEFSFileSystemRead,
		Create: resourceEFSFileSystemCreate,
		Update: resourceEFSFileSystemUpdate,
		Delete: resourceEFSFileSystemDelete,
		CustomizeDiff: resourceEFSFileSystemCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"encrypted": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"file_system_tags": {
				Type: schema.TypeSet,
				Elem: propertyFileSystemElasticFileSystemTag(),
				Optional: true,
			},
			"kms_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"lifecycle_policies": {
				Type: schema.TypeSet,
				Elem: propertyFileSystemLifecyclePolicy(),
				Optional: true,
			},
			"performance_mode": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceEFSFileSystemRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eFSFileSystemType, ResourceEFSFileSystem(), data, meta)
}

func resourceEFSFileSystemCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eFSFileSystemType, ResourceEFSFileSystem(), data, meta)
}

func resourceEFSFileSystemUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eFSFileSystemType, ResourceEFSFileSystem(), data, meta)
}

func resourceEFSFileSystemDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eFSFileSystemType, data, meta)
}

func resourceEFSFileSystemCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eFSFileSystemType, data, meta)
}
