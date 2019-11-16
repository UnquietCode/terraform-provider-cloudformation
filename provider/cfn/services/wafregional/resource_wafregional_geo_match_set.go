// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-geomatchset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalGeoMatchSetType string = "AWS::WAFRegional::GeoMatchSet"

var wAFRegionalGeoMatchSetProperties map[string]string = map[string]string{
	"geo_match_constraints": "GeoMatchConstraints",
	"name": "Name",
}

func ResourceWAFRegionalGeoMatchSet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFRegionalGeoMatchSetExists,
		Read: resourceWAFRegionalGeoMatchSetRead,
		Create: resourceWAFRegionalGeoMatchSetCreate,
		Update: resourceWAFRegionalGeoMatchSetUpdate,
		Delete: resourceWAFRegionalGeoMatchSetDelete,
		CustomizeDiff: resourceWAFRegionalGeoMatchSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"geo_match_constraints": {
				Type: schema.TypeList,
				Elem: propertyGeoMatchSetGeoMatchConstraint(),
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

func resourceWAFRegionalGeoMatchSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFRegionalGeoMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalGeoMatchSetType, ResourceWAFRegionalGeoMatchSet(), data, meta)
}

func resourceWAFRegionalGeoMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalGeoMatchSetType, ResourceWAFRegionalGeoMatchSet(), data, wAFRegionalGeoMatchSetProperties, meta)
}

func resourceWAFRegionalGeoMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalGeoMatchSetType, ResourceWAFRegionalGeoMatchSet(), data, wAFRegionalGeoMatchSetProperties, meta)
}

func resourceWAFRegionalGeoMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalGeoMatchSetType, data, meta)
}

func resourceWAFRegionalGeoMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalGeoMatchSetType, data, meta)
}
