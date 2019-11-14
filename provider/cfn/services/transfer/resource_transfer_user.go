// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceTransferUserExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceTransferUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Transfer::User", ResourceTransferUser(), data, meta)
}

func resourceTransferUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Transfer::User", ResourceTransferUser(), data, meta)
}

func resourceTransferUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Transfer::User", ResourceTransferUser(), data, meta)
}

func resourceTransferUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Transfer::User", data, meta)
}

func resourceTransferUserCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
