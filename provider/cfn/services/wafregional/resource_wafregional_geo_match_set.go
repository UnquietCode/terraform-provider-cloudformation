// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-geomatchset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

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
	return plugin.ResourceRead("AWS::WAFRegional::GeoMatchSet", ResourceWAFRegionalGeoMatchSet(), data, meta)
}

func resourceWAFRegionalGeoMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::GeoMatchSet", ResourceWAFRegionalGeoMatchSet(), data, meta)
}

func resourceWAFRegionalGeoMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::GeoMatchSet", ResourceWAFRegionalGeoMatchSet(), data, meta)
}

func resourceWAFRegionalGeoMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::GeoMatchSet", data, meta)
}

func resourceWAFRegionalGeoMatchSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

