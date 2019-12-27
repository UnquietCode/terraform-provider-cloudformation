// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html

package config

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const configOrganizationConfigRuleType string = "AWS::Config::OrganizationConfigRule"

func DatasourceConfigOrganizationConfigRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConfigOrganizationConfigRuleRead,
		
		Schema: map[string]*schema.Schema{
			"organization_managed_rule_metadata": {
				Type: schema.TypeList,
				Elem: propertyOrganizationConfigRuleOrganizationManagedRuleMetadata(),
				Optional: true,
				MaxItems: 1,
			},
			"organization_config_rule_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"organization_custom_rule_metadata": {
				Type: schema.TypeList,
				Elem: propertyOrganizationConfigRuleOrganizationCustomRuleMetadata(),
				Optional: true,
				MaxItems: 1,
			},
			"excluded_accounts": {
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

func datasourceConfigOrganizationConfigRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(configOrganizationConfigRuleType, DatasourceConfigOrganizationConfigRule(), data, meta)
}
