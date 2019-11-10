// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceXssMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceXssMatchSetCreate,
		Read:   resourceXssMatchSetRead,
		Update: resourceXssMatchSetUpdate,
		Delete: resourceXssMatchSetDelete,

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
		},
	}
}

func resourceXssMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::XssMatchSet", data, meta)
}

func resourceXssMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAF::XssMatchSet", data, meta)
}

func resourceXssMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::XssMatchSet", data, meta)
}

func resourceXssMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::XssMatchSet", data, meta)
}