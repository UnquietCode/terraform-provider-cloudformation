// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMUserToGroupAdditionType string = "AWS::IAM::UserToGroupAddition"

var iAMUserToGroupAdditionProperties map[string]string = map[string]string{
	"group_name": "GroupName",
	"users": "Users",
}

func ResourceIAMUserToGroupAddition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMUserToGroupAdditionExists,
		Read: resourceIAMUserToGroupAdditionRead,
		Create: resourceIAMUserToGroupAdditionCreate,
		Update: resourceIAMUserToGroupAdditionUpdate,
		Delete: resourceIAMUserToGroupAdditionDelete,
		CustomizeDiff: resourceIAMUserToGroupAdditionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"users": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceIAMUserToGroupAdditionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIAMUserToGroupAdditionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMUserToGroupAdditionType, ResourceIAMUserToGroupAddition(), data, meta)
}

func resourceIAMUserToGroupAdditionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(iAMUserToGroupAdditionType, ResourceIAMUserToGroupAddition(), data, iAMUserToGroupAdditionProperties, meta)
}

func resourceIAMUserToGroupAdditionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(iAMUserToGroupAdditionType, ResourceIAMUserToGroupAddition(), data, iAMUserToGroupAdditionProperties, meta)
}

func resourceIAMUserToGroupAdditionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(iAMUserToGroupAdditionType, data, meta)
}

func resourceIAMUserToGroupAdditionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(iAMUserToGroupAdditionType, data, meta)
}
