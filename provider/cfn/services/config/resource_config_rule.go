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

func ResourceConfigRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigRuleCreate,
		Read:   resourceConfigRuleRead,
		Update: resourceConfigRuleUpdate,
		Delete: resourceConfigRuleDelete,

		Schema: map[string]*schema.Schema{
			"config_rule_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"input_parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"maximum_execution_frequency": {
				Type: schema.TypeString,
				Required: false,
			},
			"scope": {
				Type: schema.TypeList,
				Elem: propertyConfigRuleScope(),
				Required: false,
				MaxItems: 1,
			},
			"source": {
				Type: schema.TypeList,
				Elem: propertyConfigRuleSource(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceConfigRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigRule", data, meta)
}

func resourceConfigRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigRule", data, meta)
}

func resourceConfigRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigRule", data, meta)
}

func resourceConfigRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigRule", data, meta)
}