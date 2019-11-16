// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMManagedPolicyType string = "AWS::IAM::ManagedPolicy"

var iAMManagedPolicyProperties map[string]string = map[string]string{
	"description": "Description",
	"groups": "Groups",
	"managed_policy_name": "ManagedPolicyName",
	"path": "Path",
	"policy_document": "PolicyDocument",
	"roles": "Roles",
	"users": "Users",
}

func ResourceIAMManagedPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMManagedPolicyExists,
		Read: resourceIAMManagedPolicyRead,
		Create: resourceIAMManagedPolicyCreate,
		Update: resourceIAMManagedPolicyUpdate,
		Delete: resourceIAMManagedPolicyDelete,
		CustomizeDiff: resourceIAMManagedPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"managed_policy_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
			"roles": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"users": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceIAMManagedPolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIAMManagedPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMManagedPolicyType, ResourceIAMManagedPolicy(), data, meta)
}

func resourceIAMManagedPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(iAMManagedPolicyType, ResourceIAMManagedPolicy(), data, iAMManagedPolicyProperties, meta)
}

func resourceIAMManagedPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(iAMManagedPolicyType, ResourceIAMManagedPolicy(), data, iAMManagedPolicyProperties, meta)
}

func resourceIAMManagedPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(iAMManagedPolicyType, data, meta)
}

func resourceIAMManagedPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(iAMManagedPolicyType, data, meta)
}
