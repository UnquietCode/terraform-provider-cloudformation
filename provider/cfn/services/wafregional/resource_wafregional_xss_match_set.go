// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-xssmatchset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalXssMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRegionalXssMatchSetCreate,
		Read:   resourceWAFRegionalXssMatchSetRead,
		Update: resourceWAFRegionalXssMatchSetUpdate,
		Delete: resourceWAFRegionalXssMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"xss_match_tuples": {
				Type: schema.TypeList,
				Elem: propertyXssMatchSetXssMatchTuple(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWAFRegionalXssMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::XssMatchSet", data, meta)
}

func resourceWAFRegionalXssMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::XssMatchSet", data, meta)
}

func resourceWAFRegionalXssMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::XssMatchSet", data, meta)
}

func resourceWAFRegionalXssMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::XssMatchSet", data, meta)
}