// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMUserCreate,
		Read:   resourceIAMUserRead,
		Update: resourceIAMUserUpdate,
		Delete: resourceIAMUserDelete,

		Schema: map[string]*schema.Schema{
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"login_profile": {
				Type: schema.TypeList,
				Elem: propertyUserLoginProfile(),
				Required: false,
				MaxItems: 1,
			},
			"managed_policy_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"path": {
				Type: schema.TypeString,
				Required: false,
			},
			"permissions_boundary": {
				Type: schema.TypeString,
				Required: false,
			},
			"policies": {
				Type: schema.TypeList,
				Elem: propertyUserPolicy(),
				Required: false,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceIAMUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::User", data, meta)
}

func resourceIAMUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::User", data, meta)
}

func resourceIAMUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::User", data, meta)
}

func resourceIAMUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::User", data, meta)
}