// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html

package iam

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const iAMManagedPolicyType string = "AWS::IAM::ManagedPolicy"

func DatasourceIAMManagedPolicy() *schema.Resource {
	return &schema.Resource{
		Read: datasourceIAMManagedPolicyRead,
		
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
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceIAMManagedPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(iAMManagedPolicyType, DatasourceIAMManagedPolicy(), data, meta)
}
