// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalSqlInjectionMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRegionalSqlInjectionMatchSetCreate,
		Read:   resourceWAFRegionalSqlInjectionMatchSetRead,
		Update: resourceWAFRegionalSqlInjectionMatchSetUpdate,
		Delete: resourceWAFRegionalSqlInjectionMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"sql_injection_match_tuples": {
				Type: schema.TypeList,
				Elem: propertySqlInjectionMatchSetSqlInjectionMatchTuple(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWAFRegionalSqlInjectionMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::SqlInjectionMatchSet", ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::SqlInjectionMatchSet", ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::SqlInjectionMatchSet", ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::SqlInjectionMatchSet", data, meta)
}