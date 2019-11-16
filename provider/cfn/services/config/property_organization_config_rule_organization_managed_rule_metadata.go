// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-organizationconfigrule-organizationmanagedrulemetadata.html

package config

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var organizationConfigRuleOrganizationManagedRuleMetadataProperties map[string]string = map[string]string{
	"tag_key_scope": "TagKeyScope",
	"tag_value_scope": "TagValueScope",
	"description": "Description",
	"resource_id_scope": "ResourceIdScope",
	"rule_identifier": "RuleIdentifier",
	"resource_types_scope": "ResourceTypesScope",
	"maximum_execution_frequency": "MaximumExecutionFrequency",
	"input_parameters": "InputParameters",
}

func propertyOrganizationConfigRuleOrganizationManagedRuleMetadata(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"tag_key_scope": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tag_value_scope": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_id_scope": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rule_identifier": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_types_scope": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"maximum_execution_frequency": {
				Type: schema.TypeString,
				Optional: true,
			},
			"input_parameters": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
