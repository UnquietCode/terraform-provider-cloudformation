// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalXssMatchSetType string = "AWS::WAFRegional::XssMatchSet"

func ResourceWAFRegionalXssMatchSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFRegionalXssMatchSetRead,
		Create: resourceWAFRegionalXssMatchSetCreate,
		Update: resourceWAFRegionalXssMatchSetUpdate,
		Delete: resourceWAFRegionalXssMatchSetDelete,
		CustomizeDiff: resourceWAFRegionalXssMatchSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"xss_match_tuples": {
				Type: schema.TypeList,
				Elem: propertyXssMatchSetXssMatchTuple(),
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

func resourceWAFRegionalXssMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalXssMatchSetType, ResourceWAFRegionalXssMatchSet(), data, meta)
}

func resourceWAFRegionalXssMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalXssMatchSetType, ResourceWAFRegionalXssMatchSet(), data, meta)
}

func resourceWAFRegionalXssMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalXssMatchSetType, ResourceWAFRegionalXssMatchSet(), data, meta)
}

func resourceWAFRegionalXssMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalXssMatchSetType, data, meta)
}

func resourceWAFRegionalXssMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalXssMatchSetType, data, meta)
}
