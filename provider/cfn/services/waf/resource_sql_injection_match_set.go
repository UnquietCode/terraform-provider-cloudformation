// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSqlInjectionMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSqlInjectionMatchSetCreate,
		Read:   resourceSqlInjectionMatchSetRead,
		Update: resourceSqlInjectionMatchSetUpdate,
		Delete: resourceSqlInjectionMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sql_injection_match_tuples": {
				Type: schema.TypeSet,
				Elem: propertySqlInjectionMatchSetSqlInjectionMatchTuple(),
				Required: false,
			},
		},
	}
}

func resourceSqlInjectionMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::SqlInjectionMatchSet", data, meta)
}

func resourceSqlInjectionMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAF::SqlInjectionMatchSet", data, meta)
}

func resourceSqlInjectionMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::SqlInjectionMatchSet", data, meta)
}

func resourceSqlInjectionMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::SqlInjectionMatchSet", data, meta)
}