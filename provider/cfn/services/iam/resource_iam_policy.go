// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMPolicyType string = "AWS::IAM::Policy"

var iAMPolicyProperties map[string]string = map[string]string{
	"groups": "Groups",
	"policy_document": "PolicyDocument",
	"policy_name": "PolicyName",
	"roles": "Roles",
	"users": "Users",
}

func ResourceIAMPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMPolicyExists,
		Read: resourceIAMPolicyRead,
		Create: resourceIAMPolicyCreate,
		Update: resourceIAMPolicyUpdate,
		Delete: resourceIAMPolicyDelete,
		CustomizeDiff: resourceIAMPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
			"policy_name": {
				Type: schema.TypeString,
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

func resourceIAMPolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIAMPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMPolicyType, ResourceIAMPolicy(), data, meta)
}

func resourceIAMPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(iAMPolicyType, ResourceIAMPolicy(), data, iAMPolicyProperties, meta)
}

func resourceIAMPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(iAMPolicyType, ResourceIAMPolicy(), data, iAMPolicyProperties, meta)
}

func resourceIAMPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(iAMPolicyType, data, meta)
}

func resourceIAMPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(iAMPolicyType, data, meta)
}
