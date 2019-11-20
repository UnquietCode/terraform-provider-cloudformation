// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFIPSetType string = "AWS::WAF::IPSet"

func ResourceWAFIPSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFIPSetRead,
		Create: resourceWAFIPSetCreate,
		Update: resourceWAFIPSetUpdate,
		Delete: resourceWAFIPSetDelete,
		CustomizeDiff: resourceWAFIPSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"ip_set_descriptors": {
				Type: schema.TypeSet,
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

func resourceWAFIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFIPSetType, ResourceWAFIPSet(), data, meta)
}

func resourceWAFIPSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFIPSetType, ResourceWAFIPSet(), data, meta)
}

func resourceWAFIPSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFIPSetType, ResourceWAFIPSet(), data, meta)
}

func resourceWAFIPSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFIPSetType, data, meta)
}

func resourceWAFIPSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFIPSetType, data, meta)
}
