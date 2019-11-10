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

func ResourceWAFRegionalRegexPatternSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRegionalRegexPatternSetCreate,
		Read:   resourceWAFRegionalRegexPatternSetRead,
		Update: resourceWAFRegionalRegexPatternSetUpdate,
		Delete: resourceWAFRegionalRegexPatternSetDelete,

		Schema: map[string]*schema.Schema{
			"regex_pattern_strings": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWAFRegionalRegexPatternSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::RegexPatternSet", data, meta)
}

func resourceWAFRegionalRegexPatternSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::RegexPatternSet", data, meta)
}

func resourceWAFRegionalRegexPatternSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::RegexPatternSet", data, meta)
}

func resourceWAFRegionalRegexPatternSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::RegexPatternSet", data, meta)
}