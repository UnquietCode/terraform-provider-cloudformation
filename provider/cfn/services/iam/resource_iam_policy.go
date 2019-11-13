// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIAMPolicyExists,
		Read:   resourceIAMPolicyRead,
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
	return plugin.ResourceRead("AWS::IAM::Policy", ResourceIAMPolicy(), data, meta)
}

func resourceIAMPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::Policy", ResourceIAMPolicy(), data, meta)
}

func resourceIAMPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::Policy", ResourceIAMPolicy(), data, meta)
}

func resourceIAMPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::Policy", data, meta)
}

func resourceIAMPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}