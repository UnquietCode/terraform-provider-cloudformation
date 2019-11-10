// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconfigrule.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigOrganizationConfigRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigOrganizationConfigRuleCreate,
		Read:   resourceConfigOrganizationConfigRuleRead,
		Update: resourceConfigOrganizationConfigRuleUpdate,
		Delete: resourceConfigOrganizationConfigRuleDelete,

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
				ForceNew: true,
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
		},
	}
}

func resourceConfigOrganizationConfigRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::OrganizationConfigRule", data, meta)
}

func resourceConfigOrganizationConfigRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::OrganizationConfigRule", data, meta)
}

func resourceConfigOrganizationConfigRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::OrganizationConfigRule", data, meta)
}

func resourceConfigOrganizationConfigRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::OrganizationConfigRule", data, meta)
}