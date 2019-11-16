// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ipset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalIPSetType string = "AWS::WAFRegional::IPSet"

func ResourceWAFRegionalIPSet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFRegionalIPSetExists,
		Read: resourceWAFRegionalIPSetRead,
		Create: resourceWAFRegionalIPSetCreate,
		Update: resourceWAFRegionalIPSetUpdate,
		Delete: resourceWAFRegionalIPSetDelete,
		CustomizeDiff: resourceWAFRegionalIPSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"ip_set_descriptors": {
				Type: schema.TypeList,
				Elem: propertyIPSetIPSetDescriptor(),
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

func resourceWAFRegionalIPSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFRegionalIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalIPSetType, ResourceWAFRegionalIPSet(), data, meta)
}

func resourceWAFRegionalIPSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalIPSetType, ResourceWAFRegionalIPSet(), data, meta)
}

func resourceWAFRegionalIPSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalIPSetType, ResourceWAFRegionalIPSet(), data, meta)
}

func resourceWAFRegionalIPSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalIPSetType, data, meta)
}

func resourceWAFRegionalIPSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalIPSetType, data, meta)
}
