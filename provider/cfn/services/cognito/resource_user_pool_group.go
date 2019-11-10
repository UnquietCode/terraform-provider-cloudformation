// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceUserPoolGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPoolGroupCreate,
		Read:   resourceUserPoolGroupRead,
		Update: resourceUserPoolGroupUpdate,
		Delete: resourceUserPoolGroupDelete,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"precedence": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceUserPoolGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Cognito::UserPoolGroup", data, meta)
}

func resourceUserPoolGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Cognito::UserPoolGroup", data, meta)
}

func resourceUserPoolGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Cognito::UserPoolGroup", data, meta)
}

func resourceUserPoolGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Cognito::UserPoolGroup", data, meta)
}