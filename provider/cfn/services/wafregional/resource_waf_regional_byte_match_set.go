// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalByteMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRegionalByteMatchSetCreate,
		Read:   resourceWAFRegionalByteMatchSetRead,
		Update: resourceWAFRegionalByteMatchSetUpdate,
		Delete: resourceWAFRegionalByteMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"byte_match_tuples": {
				Type: schema.TypeList,
				Elem: propertyByteMatchTuple(),
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

func resourceWAFRegionalByteMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::ByteMatchSet", data, meta)
}

func resourceWAFRegionalByteMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::ByteMatchSet", data, meta)
}

func resourceWAFRegionalByteMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::ByteMatchSet", data, meta)
}

func resourceWAFRegionalByteMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::ByteMatchSet", data, meta)
}