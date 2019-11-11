// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFByteMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFByteMatchSetCreate,
		Read:   resourceWAFByteMatchSetRead,
		Update: resourceWAFByteMatchSetUpdate,
		Delete: resourceWAFByteMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"byte_match_tuples": {
				Type: schema.TypeSet,
				Elem: propertyByteMatchSetByteMatchTuple(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceWAFByteMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::ByteMatchSet", ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAF::ByteMatchSet", ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::ByteMatchSet", ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::ByteMatchSet", data, meta)
}