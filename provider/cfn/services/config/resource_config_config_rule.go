// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigConfigRule() *schema.Resource {
	return &schema.Resource{
		Exists: resourceConfigConfigRuleExists,
		Read: resourceConfigConfigRuleRead,
		Create: resourceConfigConfigRuleCreate,
		Update: resourceConfigConfigRuleUpdate,
		Delete: resourceConfigConfigRuleDelete,
		CustomizeDiff: resourceConfigConfigRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"config_rule_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"input_parameters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"maximum_execution_frequency": {
				Type: schema.TypeString,
				Optional: true,
			},
			"scope": {
				Type: schema.TypeList,
				Elem: propertyConfigRuleScope(),
				Optional: true,
				MaxItems: 1,
			},
			"source": {
				Type: schema.TypeList,
				Elem: propertyConfigRuleSource(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConfigConfigRuleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceConfigConfigRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigRule", ResourceConfigConfigRule(), data, meta)
}

func resourceConfigConfigRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigRule", ResourceConfigConfigRule(), data, meta)
}

func resourceConfigConfigRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigRule", ResourceConfigConfigRule(), data, meta)
}

func resourceConfigConfigRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigRule", data, meta)
}

func resourceConfigConfigRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::Config::ConfigRule", data, meta)
}

