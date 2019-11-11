// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIAMManagedPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMManagedPolicyCreate,
		Read:   resourceIAMManagedPolicyRead,
		Update: resourceIAMManagedPolicyUpdate,
		Delete: resourceIAMManagedPolicyDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"path": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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

func resourceIAMManagedPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::ManagedPolicy", ResourceIAMManagedPolicy(), data, meta)
}

func resourceIAMManagedPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::ManagedPolicy", ResourceIAMManagedPolicy(), data, meta)
}

func resourceIAMManagedPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::ManagedPolicy", ResourceIAMManagedPolicy(), data, meta)
}

func resourceIAMManagedPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::ManagedPolicy", data, meta)
}