// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMUserType string = "AWS::IAM::User"

func ResourceIAMUser() *schema.Resource {
	return &schema.Resource{
		Read: resourceIAMUserRead,
		Create: resourceIAMUserCreate,
		Update: resourceIAMUserUpdate,
		Delete: resourceIAMUserDelete,
		CustomizeDiff: resourceIAMUserCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"login_profile": {
				Type: schema.TypeSet,
				Elem: propertyUserLoginProfile(),
				Optional: true,
				MaxItems: 1,
			},
			"managed_policy_arns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"permissions_boundary": {
				Type: schema.TypeString,
				Optional: true,
			},
			"policies": {
				Type: schema.TypeList,
				Elem: propertyUserPolicy(),
				Optional: true,
			},
			"user_name": {
				Type: schema.TypeString,
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

func resourceIAMUserRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMUserType, ResourceIAMUser(), data, meta)
}

func resourceIAMUserCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(iAMUserType, ResourceIAMUser(), data, meta)
}

func resourceIAMUserUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(iAMUserType, ResourceIAMUser(), data, meta)
}

func resourceIAMUserDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(iAMUserType, data, meta)
}

func resourceIAMUserCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(iAMUserType, data, meta)
}
