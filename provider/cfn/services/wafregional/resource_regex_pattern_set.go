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

func ResourceRegexPatternSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceRegexPatternSetCreate,
		Read:   resourceRegexPatternSetRead,
		Update: resourceRegexPatternSetUpdate,
		Delete: resourceRegexPatternSetDelete,

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

func resourceRegexPatternSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::RegexPatternSet", data, meta)
}

func resourceRegexPatternSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::RegexPatternSet", data, meta)
}

func resourceRegexPatternSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::RegexPatternSet", data, meta)
}

func resourceRegexPatternSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::RegexPatternSet", data, meta)
}