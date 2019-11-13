// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceWAFRegionalSqlInjectionMatchSetExists,
		Read:   resourceWAFRegionalSqlInjectionMatchSetRead,
		Create: resourceWAFRegionalSqlInjectionMatchSetCreate,
		Update: resourceWAFRegionalSqlInjectionMatchSetUpdate,
		Delete: resourceWAFRegionalSqlInjectionMatchSetDelete,
		CustomizeDiff: resourceWAFRegionalSqlInjectionMatchSetCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"sql_injection_match_tuples": {
				Type: schema.TypeList,
				Elem: propertySqlInjectionMatchSetSqlInjectionMatchTuple(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWAFRegionalSqlInjectionMatchSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::SqlInjectionMatchSet", ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::SqlInjectionMatchSet", ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::SqlInjectionMatchSet", ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::SqlInjectionMatchSet", data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}