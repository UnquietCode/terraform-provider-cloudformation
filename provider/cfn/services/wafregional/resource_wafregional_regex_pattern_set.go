// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalRegexPatternSetType string = "AWS::WAFRegional::RegexPatternSet"

func ResourceWAFRegionalRegexPatternSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFRegionalRegexPatternSetRead,
		Create: resourceWAFRegionalRegexPatternSetCreate,
		Update: resourceWAFRegionalRegexPatternSetUpdate,
		Delete: resourceWAFRegionalRegexPatternSetDelete,
		CustomizeDiff: resourceWAFRegionalRegexPatternSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"regex_pattern_strings": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
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

func resourceWAFRegionalRegexPatternSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalRegexPatternSetType, ResourceWAFRegionalRegexPatternSet(), data, meta)
}

func resourceWAFRegionalRegexPatternSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalRegexPatternSetType, ResourceWAFRegionalRegexPatternSet(), data, meta)
}

func resourceWAFRegionalRegexPatternSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalRegexPatternSetType, ResourceWAFRegionalRegexPatternSet(), data, meta)
}

func resourceWAFRegionalRegexPatternSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalRegexPatternSetType, data, meta)
}

func resourceWAFRegionalRegexPatternSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalRegexPatternSetType, data, meta)
}
