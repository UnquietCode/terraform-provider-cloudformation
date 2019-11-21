// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sqlinjectionmatchset.html

package wafregional

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalSqlInjectionMatchSetType string = "AWS::WAFRegional::SqlInjectionMatchSet"

func ResourceWAFRegionalSqlInjectionMatchSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFRegionalSqlInjectionMatchSetRead,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceWAFRegionalSqlInjectionMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalSqlInjectionMatchSetType, ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalSqlInjectionMatchSetType, ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalSqlInjectionMatchSetType, ResourceWAFRegionalSqlInjectionMatchSet(), data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalSqlInjectionMatchSetType, data, meta)
}

func resourceWAFRegionalSqlInjectionMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalSqlInjectionMatchSetType, data, meta)
}
