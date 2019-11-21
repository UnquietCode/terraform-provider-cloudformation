// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html

package wafregional

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalWebACLType string = "AWS::WAFRegional::WebACL"

func ResourceWAFRegionalWebACL() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFRegionalWebACLRead,
		Create: resourceWAFRegionalWebACLCreate,
		Update: resourceWAFRegionalWebACLUpdate,
		Delete: resourceWAFRegionalWebACLDelete,
		CustomizeDiff: resourceWAFRegionalWebACLCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"metric_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_action": {
				Type: schema.TypeSet,
				Elem: propertyWebACLAction(),
				Required: true,
				MaxItems: 1,
			},
			"rules": {
				Type: schema.TypeList,
				Elem: propertyWebACLRule(),
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

func resourceWAFRegionalWebACLRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalWebACLType, ResourceWAFRegionalWebACL(), data, meta)
}

func resourceWAFRegionalWebACLCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalWebACLType, ResourceWAFRegionalWebACL(), data, meta)
}

func resourceWAFRegionalWebACLUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalWebACLType, ResourceWAFRegionalWebACL(), data, meta)
}

func resourceWAFRegionalWebACLDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalWebACLType, data, meta)
}

func resourceWAFRegionalWebACLCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalWebACLType, data, meta)
}
