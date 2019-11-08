// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRuleCreate,
		Read:   resourceWAFRuleRead,
		Update: resourceWAFRuleUpdate,
		Delete: resourceWAFRuleDelete,

		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"predicates": {
				Type: schema.TypeSet,
				Elem: propertyPredicate(),
				Required: false,
			},
		},
	}
}

func resourceWAFRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::Rule", data, meta)
}

func resourceWAFRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAF::Rule", data, meta)
}

func resourceWAFRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::Rule", data, meta)
}

func resourceWAFRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::Rule", data, meta)
}