// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appstream

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDirectoryConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceDirectoryConfigCreate,
		Read:   resourceDirectoryConfigRead,
		Update: resourceDirectoryConfigUpdate,
		Delete: resourceDirectoryConfigDelete,

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
				ForceNew: true,
			},
		},
	}
}

func resourceDirectoryConfigCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AppStream::DirectoryConfig", data, meta)
}

func resourceDirectoryConfigRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AppStream::DirectoryConfig", data, meta)
}

func resourceDirectoryConfigUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AppStream::DirectoryConfig", data, meta)
}

func resourceDirectoryConfigDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AppStream::DirectoryConfig", data, meta)
}