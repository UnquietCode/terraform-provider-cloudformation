// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-regexpatternset.html

package wafv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFv2RegexPatternSetType string = "AWS::WAFv2::RegexPatternSet"

func DatasourceWAFv2RegexPatternSet() *schema.Resource {
	return &schema.Resource{
		Read: datasourceWAFv2RegexPatternSetRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"regular_expression_list": {
				Type: schema.TypeList,
				Elem: propertyRegexPatternSetRegularExpressionList(),
				Optional: true,
				MaxItems: 1,
			},
			"scope": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyRegexPatternSetTagList(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceWAFv2RegexPatternSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFv2RegexPatternSetType, DatasourceWAFv2RegexPatternSet(), data, meta)
}
