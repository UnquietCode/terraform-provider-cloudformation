// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceFSxFileSystem() *schema.Resource {
	return &schema.Resource{
		Read: datasourceFSxFileSystemRead,
		
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
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceFSxFileSystemRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(fSxFileSystemType, DatasourceFSxFileSystem(), data, meta)
}
