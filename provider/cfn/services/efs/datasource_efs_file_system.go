// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html

package efs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eFSFileSystemType string = "AWS::EFS::FileSystem"

func DatasourceEFSFileSystem() *schema.Resource {
	return &schema.Resource{
		Read: datasourceEFSFileSystemRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceEFSFileSystemRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eFSFileSystemType, DatasourceEFSFileSystem(), data, meta)
}
