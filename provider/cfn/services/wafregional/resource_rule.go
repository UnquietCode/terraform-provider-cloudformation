// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRuleCreate,
		Read:   resourceRuleRead,
		Update: resourceRuleUpdate,
		Delete: resourceRuleDelete,

		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"predicates": {
				Type: schema.TypeList,
				Elem: propertyRulePredicate(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::Rule", data, meta)
}

func resourceRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::Rule", data, meta)
}

func resourceRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::Rule", data, meta)
}

func resourceRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::Rule", data, meta)
}