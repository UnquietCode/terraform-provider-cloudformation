// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMUserToGroupAddition() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMUserToGroupAdditionCreate,
		Read:   resourceIAMUserToGroupAdditionRead,
		Update: resourceIAMUserToGroupAdditionUpdate,
		Delete: resourceIAMUserToGroupAdditionDelete,

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
		},
	}
}

func resourceIAMUserToGroupAdditionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::UserToGroupAddition", data, meta)
}

func resourceIAMUserToGroupAdditionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::UserToGroupAddition", data, meta)
}

func resourceIAMUserToGroupAdditionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::UserToGroupAddition", data, meta)
}

func resourceIAMUserToGroupAdditionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::UserToGroupAddition", data, meta)
}