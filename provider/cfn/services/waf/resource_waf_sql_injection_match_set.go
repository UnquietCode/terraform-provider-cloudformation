// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html

package waf

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFSqlInjectionMatchSetType string = "AWS::WAF::SqlInjectionMatchSet"

func ResourceWAFSqlInjectionMatchSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFSqlInjectionMatchSetRead,
		Create: resourceWAFSqlInjectionMatchSetCreate,
		Update: resourceWAFSqlInjectionMatchSetUpdate,
		Delete: resourceWAFSqlInjectionMatchSetDelete,
		CustomizeDiff: resourceWAFSqlInjectionMatchSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"sql_injection_match_tuples": {
				Type: schema.TypeSet,
				Elem: propertySqlInjectionMatchSetSqlInjectionMatchTuple(),
				Optional: true,
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

func resourceWAFSqlInjectionMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFSqlInjectionMatchSetType, ResourceWAFSqlInjectionMatchSet(), data, meta)
}

func resourceWAFSqlInjectionMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFSqlInjectionMatchSetType, ResourceWAFSqlInjectionMatchSet(), data, meta)
}

func resourceWAFSqlInjectionMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFSqlInjectionMatchSetType, ResourceWAFSqlInjectionMatchSet(), data, meta)
}

func resourceWAFSqlInjectionMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFSqlInjectionMatchSetType, data, meta)
}

func resourceWAFSqlInjectionMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFSqlInjectionMatchSetType, data, meta)
}
