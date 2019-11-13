// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceWAFByteMatchSetExists,
		Read:   resourceWAFByteMatchSetRead,
		Create: resourceWAFByteMatchSetCreate,
		Update: resourceWAFByteMatchSetUpdate,
		Delete: resourceWAFByteMatchSetDelete,
		CustomizeDiff: resourceWAFByteMatchSetCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"byte_match_tuples": {
				Type: schema.TypeSet,
				Elem: propertyByteMatchSetByteMatchTuple(),
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

func resourceWAFByteMatchSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFByteMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAF::ByteMatchSet", ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::ByteMatchSet", ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::ByteMatchSet", ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::ByteMatchSet", data, meta)
}

func resourceWAFByteMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}