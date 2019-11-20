// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-bytematchset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalByteMatchSetType string = "AWS::WAFRegional::ByteMatchSet"

func ResourceWAFRegionalByteMatchSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFRegionalByteMatchSetRead,
		Create: resourceWAFRegionalByteMatchSetCreate,
		Update: resourceWAFRegionalByteMatchSetUpdate,
		Delete: resourceWAFRegionalByteMatchSetDelete,
		CustomizeDiff: resourceWAFRegionalByteMatchSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"byte_match_tuples": {
				Type: schema.TypeList,
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

func resourceWAFRegionalByteMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalByteMatchSetType, ResourceWAFRegionalByteMatchSet(), data, meta)
}

func resourceWAFRegionalByteMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalByteMatchSetType, ResourceWAFRegionalByteMatchSet(), data, meta)
}

func resourceWAFRegionalByteMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalByteMatchSetType, ResourceWAFRegionalByteMatchSet(), data, meta)
}

func resourceWAFRegionalByteMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalByteMatchSetType, data, meta)
}

func resourceWAFRegionalByteMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalByteMatchSetType, data, meta)
}
