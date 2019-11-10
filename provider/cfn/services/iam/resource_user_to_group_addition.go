// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceUserToGroupAddition() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserToGroupAdditionCreate,
		Read:   resourceUserToGroupAdditionRead,
		Update: resourceUserToGroupAdditionUpdate,
		Delete: resourceUserToGroupAdditionDelete,

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

func resourceUserToGroupAdditionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::UserToGroupAddition", data, meta)
}

func resourceUserToGroupAdditionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::UserToGroupAddition", data, meta)
}

func resourceUserToGroupAdditionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::UserToGroupAddition", data, meta)
}

func resourceUserToGroupAdditionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::UserToGroupAddition", data, meta)
}