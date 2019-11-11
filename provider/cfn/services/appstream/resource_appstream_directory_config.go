// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAppStreamDirectoryConfig() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAppStreamDirectoryConfigExists,
		Read: resourceAppStreamDirectoryConfigRead,
		Create: resourceAppStreamDirectoryConfigCreate,
		Update: resourceAppStreamDirectoryConfigUpdate,
		Delete: resourceAppStreamDirectoryConfigDelete,
		CustomizeDiff: resourceAppStreamDirectoryConfigCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"organizational_unit_distinguished_names": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"service_account_credentials": {
				Type: schema.TypeList,
				Elem: propertyDirectoryConfigServiceAccountCredentials(),
				Required: true,
				MaxItems: 1,
			},
			"directory_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAppStreamDirectoryConfigExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAppStreamDirectoryConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::DirectoryConfig", ResourceAppStreamDirectoryConfig(), data, meta)
}

func resourceAppStreamDirectoryConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::DirectoryConfig", ResourceAppStreamDirectoryConfig(), data, meta)
}

func resourceAppStreamDirectoryConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::DirectoryConfig", ResourceAppStreamDirectoryConfig(), data, meta)
}

func resourceAppStreamDirectoryConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::DirectoryConfig", data, meta)
}

func resourceAppStreamDirectoryConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::AppStream::DirectoryConfig", data, meta)
}

