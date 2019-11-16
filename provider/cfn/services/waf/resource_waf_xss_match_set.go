// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFXssMatchSetType string = "AWS::WAF::XssMatchSet"

var wAFXssMatchSetProperties map[string]string = map[string]string{
	"name": "Name",
	"xss_match_tuples": "XssMatchTuples",
}

func ResourceWAFXssMatchSet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFXssMatchSetExists,
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
			},
		},
	}
}

func resourceWAFXssMatchSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFXssMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFXssMatchSetType, ResourceWAFXssMatchSet(), data, meta)
}

func resourceWAFXssMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFXssMatchSetType, ResourceWAFXssMatchSet(), data, wAFXssMatchSetProperties, meta)
}

func resourceWAFXssMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFXssMatchSetType, ResourceWAFXssMatchSet(), data, wAFXssMatchSetProperties, meta)
}

func resourceWAFXssMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFXssMatchSetType, data, meta)
}

func resourceWAFXssMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFXssMatchSetType, data, meta)
}
