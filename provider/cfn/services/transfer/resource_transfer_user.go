// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-user.html

package transfer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const transferUserType string = "AWS::Transfer::User"

func ResourceTransferUser() *schema.Resource {
	return &schema.Resource{
		Exists: resourceTransferUserExists,
		Read: resourceTransferUserRead,
		Create: resourceTransferUserCreate,
		Update: resourceTransferUserUpdate,
		Delete: resourceTransferUserDelete,
		CustomizeDiff: resourceTransferUserCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"policy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role": {
				Type: schema.TypeString,
				Required: true,
			},
			"home_directory": {
				Type: schema.TypeString,
				Optional: true,
			},
			"server_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ssh_public_keys": {
				Type: schema.TypeList,
				Elem: propertyUserSshPublicKey(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceTransferUserExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceTransferUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(transferUserType, ResourceTransferUser(), data, meta)
}

func resourceTransferUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(transferUserType, ResourceTransferUser(), data, meta)
}

func resourceTransferUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(transferUserType, ResourceTransferUser(), data, meta)
}

func resourceTransferUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(transferUserType, data, meta)
}

func resourceTransferUserCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(transferUserType, data, meta)
}
