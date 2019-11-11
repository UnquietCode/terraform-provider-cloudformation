// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFXssMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFXssMatchSetCreate,
		Read:   resourceWAFXssMatchSetRead,
		Update: resourceWAFXssMatchSetUpdate,
		Delete: resourceWAFXssMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourceWAFXssMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::XssMatchSet", ResourceWAFXssMatchSet(), data, meta)
}

func resourceWAFXssMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAF::XssMatchSet", ResourceWAFXssMatchSet(), data, meta)
}

func resourceWAFXssMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::XssMatchSet", ResourceWAFXssMatchSet(), data, meta)
}

func resourceWAFXssMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::XssMatchSet", data, meta)
}