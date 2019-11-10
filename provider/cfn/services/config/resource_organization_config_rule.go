// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceOrganizationConfigRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceOrganizationConfigRuleCreate,
		Read:   resourceOrganizationConfigRuleRead,
		Update: resourceOrganizationConfigRuleUpdate,
		Delete: resourceOrganizationConfigRuleDelete,

		Schema: map[string]*schema.Schema{
			"organization_managed_rule_metadata": {
				Type: schema.TypeList,
				Elem: propertyOrganizationConfigRuleOrganizationManagedRuleMetadata(),
				Required: false,
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
				Required: false,
				MaxItems: 1,
			},
			"excluded_accounts": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}

func resourceOrganizationConfigRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::OrganizationConfigRule", data, meta)
}

func resourceOrganizationConfigRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::OrganizationConfigRule", data, meta)
}

func resourceOrganizationConfigRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::OrganizationConfigRule", data, meta)
}

func resourceOrganizationConfigRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::OrganizationConfigRule", data, meta)
}