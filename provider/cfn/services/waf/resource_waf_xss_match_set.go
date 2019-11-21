// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html

package waf

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFXssMatchSetType string = "AWS::WAF::XssMatchSet"

func ResourceWAFXssMatchSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFXssMatchSetRead,
		Create: resourceWAFXssMatchSetCreate,
		Update: resourceWAFXssMatchSetUpdate,
		Delete: resourceWAFXssMatchSetDelete,
		CustomizeDiff: resourceWAFXssMatchSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"xss_match_tuples": {
				Type: schema.TypeSet,
				Elem: propertyXssMatchSetXssMatchTuple(),
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

func resourceWAFXssMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFXssMatchSetType, ResourceWAFXssMatchSet(), data, meta)
}

func resourceWAFXssMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFXssMatchSetType, ResourceWAFXssMatchSet(), data, meta)
}

func resourceWAFXssMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFXssMatchSetType, ResourceWAFXssMatchSet(), data, meta)
}

func resourceWAFXssMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFXssMatchSetType, data, meta)
}

func resourceWAFXssMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFXssMatchSetType, data, meta)
}
