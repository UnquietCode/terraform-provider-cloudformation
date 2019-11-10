// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceByteMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceByteMatchSetCreate,
		Read:   resourceByteMatchSetRead,
		Update: resourceByteMatchSetUpdate,
		Delete: resourceByteMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"byte_match_tuples": {
				Type: schema.TypeList,
				Elem: propertyByteMatchSetByteMatchTuple(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceByteMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::ByteMatchSet", data, meta)
}

func resourceByteMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::ByteMatchSet", data, meta)
}

func resourceByteMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::ByteMatchSet", data, meta)
}

func resourceByteMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::ByteMatchSet", data, meta)
}