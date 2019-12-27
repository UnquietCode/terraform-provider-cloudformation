// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html

package wafv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFv2WebACLType string = "AWS::WAFv2::WebACL"

func DatasourceWAFv2WebACL() *schema.Resource {
	return &schema.Resource{
		Read: datasourceWAFv2WebACLRead,
		
		Schema: map[string]*schema.Schema{
			"default_action": {
				Type: schema.TypeList,
				Elem: propertyWebACLDefaultAction(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"scope": {
				Type: schema.TypeString,
				Required: true,
			},
			"rules": {
				Type: schema.TypeList,
				Elem: propertyWebACLRules(),
				Optional: true,
				MaxItems: 1,
			},
			"visibility_config": {
				Type: schema.TypeList,
				Elem: propertyWebACLVisibilityConfig(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyWebACLTagList(),
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

func datasourceWAFv2WebACLRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFv2WebACLType, DatasourceWAFv2WebACL(), data, meta)
}
