// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package directoryservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDirectoryServiceMicrosoftAD() *schema.Resource {
	return &schema.Resource{
		Create: resourceDirectoryServiceMicrosoftADCreate,
		Read:   resourceDirectoryServiceMicrosoftADRead,
		Update: resourceDirectoryServiceMicrosoftADUpdate,
		Delete: resourceDirectoryServiceMicrosoftADDelete,

		Schema: map[string]*schema.Schema{
			"create_alias": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"edition": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"enable_sso": {
				Type: schema.TypeBool,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"short_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"vpc_settings": {
				Type: schema.TypeList,
				Elem: propertyVpcSettings(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceDirectoryServiceMicrosoftADCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DirectoryService::MicrosoftAD", data, meta)
}

func resourceDirectoryServiceMicrosoftADRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DirectoryService::MicrosoftAD", data, meta)
}

func resourceDirectoryServiceMicrosoftADUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DirectoryService::MicrosoftAD", data, meta)
}

func resourceDirectoryServiceMicrosoftADDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DirectoryService::MicrosoftAD", data, meta)
}