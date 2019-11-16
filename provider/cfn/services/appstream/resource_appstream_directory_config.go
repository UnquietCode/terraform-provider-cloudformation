// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const appStreamDirectoryConfigType string = "AWS::AppStream::DirectoryConfig"

var appStreamDirectoryConfigProperties map[string]string = map[string]string{
	"organizational_unit_distinguished_names": "OrganizationalUnitDistinguishedNames",
	"service_account_credentials": "ServiceAccountCredentials",
	"directory_name": "DirectoryName",
}

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
				Type: schema.TypeSet,
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
	return plugin.ResourceRead(appStreamDirectoryConfigType, ResourceAppStreamDirectoryConfig(), data, meta)
}

func resourceAppStreamDirectoryConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(appStreamDirectoryConfigType, ResourceAppStreamDirectoryConfig(), data, appStreamDirectoryConfigProperties, meta)
}

func resourceAppStreamDirectoryConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(appStreamDirectoryConfigType, ResourceAppStreamDirectoryConfig(), data, appStreamDirectoryConfigProperties, meta)
}

func resourceAppStreamDirectoryConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(appStreamDirectoryConfigType, data, meta)
}

func resourceAppStreamDirectoryConfigCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(appStreamDirectoryConfigType, data, meta)
}
