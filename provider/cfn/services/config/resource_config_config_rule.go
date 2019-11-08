// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigConfigRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigConfigRuleCreate,
		Read:   resourceConfigConfigRuleRead,
		Update: resourceConfigConfigRuleUpdate,
		Delete: resourceConfigConfigRuleDelete,

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
				Elem: propertyScope(),
				Required: false,
				MaxItems: 1,
			},
			"source": {
				Type: schema.TypeList,
				Elem: propertySource(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceConfigConfigRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigRule", data, meta)
}

func resourceConfigConfigRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigRule", data, meta)
}

func resourceConfigConfigRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigRule", data, meta)
}

func resourceConfigConfigRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigRule", data, meta)
}