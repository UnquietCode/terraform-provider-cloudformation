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

func ResourceWAFRegionalGeoMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRegionalGeoMatchSetCreate,
		Read:   resourceWAFRegionalGeoMatchSetRead,
		Update: resourceWAFRegionalGeoMatchSetUpdate,
		Delete: resourceWAFRegionalGeoMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"geo_match_constraints": {
				Type: schema.TypeList,
				Elem: propertyGeoMatchConstraint(),
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

func resourceWAFRegionalGeoMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::GeoMatchSet", data, meta)
}

func resourceWAFRegionalGeoMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::GeoMatchSet", data, meta)
}

func resourceWAFRegionalGeoMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::GeoMatchSet", data, meta)
}

func resourceWAFRegionalGeoMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::GeoMatchSet", data, meta)
}