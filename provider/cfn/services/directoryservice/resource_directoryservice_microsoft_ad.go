// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html

package directoryservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const directoryServiceMicrosoftADType string = "AWS::DirectoryService::MicrosoftAD"

var directoryServiceMicrosoftADProperties map[string]string = map[string]string{
	"create_alias": "CreateAlias",
	"edition": "Edition",
	"enable_sso": "EnableSso",
	"name": "Name",
	"password": "Password",
	"short_name": "ShortName",
	"vpc_settings": "VpcSettings",
}

func ResourceDirectoryServiceMicrosoftAD() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDirectoryServiceMicrosoftADExists,
		Read: resourceDirectoryServiceMicrosoftADRead,
		Create: resourceDirectoryServiceMicrosoftADCreate,
		Update: resourceDirectoryServiceMicrosoftADUpdate,
		Delete: resourceDirectoryServiceMicrosoftADDelete,
		CustomizeDiff: resourceDirectoryServiceMicrosoftADCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"create_alias": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"edition": {
				Type: schema.TypeString,
				Optional: true,
			},
			"enable_sso": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"password": {
				Type: schema.TypeString,
				Required: true,
			},
			"short_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"vpc_settings": {
				Type: schema.TypeList,
				Elem: propertyMicrosoftADVpcSettings(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDirectoryServiceMicrosoftADExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDirectoryServiceMicrosoftADRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(directoryServiceMicrosoftADType, ResourceDirectoryServiceMicrosoftAD(), data, meta)
}

func resourceDirectoryServiceMicrosoftADCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(directoryServiceMicrosoftADType, ResourceDirectoryServiceMicrosoftAD(), data, directoryServiceMicrosoftADProperties, meta)
}

func resourceDirectoryServiceMicrosoftADUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(directoryServiceMicrosoftADType, ResourceDirectoryServiceMicrosoftAD(), data, directoryServiceMicrosoftADProperties, meta)
}

func resourceDirectoryServiceMicrosoftADDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(directoryServiceMicrosoftADType, data, meta)
}

func resourceDirectoryServiceMicrosoftADCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(directoryServiceMicrosoftADType, data, meta)
}
