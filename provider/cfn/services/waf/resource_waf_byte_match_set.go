// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html

package waf

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFByteMatchSetType string = "AWS::WAF::ByteMatchSet"

func ResourceWAFByteMatchSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFByteMatchSetRead,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceWAFByteMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFByteMatchSetType, ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFByteMatchSetType, ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFByteMatchSetType, ResourceWAFByteMatchSet(), data, meta)
}

func resourceWAFByteMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFByteMatchSetType, data, meta)
}

func resourceWAFByteMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFByteMatchSetType, data, meta)
}
