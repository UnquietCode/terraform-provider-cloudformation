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

func ResourceIAMManagedPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMManagedPolicyCreate,
		Read:   resourceIAMManagedPolicyRead,
		Update: resourceIAMManagedPolicyUpdate,
		Delete: resourceIAMManagedPolicyDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"managed_policy_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"path": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
			"roles": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"users": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceIAMManagedPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IAM::ManagedPolicy", data, meta)
}

func resourceIAMManagedPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IAM::ManagedPolicy", data, meta)
}

func resourceIAMManagedPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IAM::ManagedPolicy", data, meta)
}

func resourceIAMManagedPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IAM::ManagedPolicy", data, meta)
}