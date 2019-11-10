// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGeoMatchSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceGeoMatchSetCreate,
		Read:   resourceGeoMatchSetRead,
		Update: resourceGeoMatchSetUpdate,
		Delete: resourceGeoMatchSetDelete,

		Schema: map[string]*schema.Schema{
			"geo_match_constraints": {
				Type: schema.TypeList,
				Elem: propertyGeoMatchSetGeoMatchConstraint(),
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

func resourceGeoMatchSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::GeoMatchSet", data, meta)
}

func resourceGeoMatchSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::GeoMatchSet", data, meta)
}

func resourceGeoMatchSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::GeoMatchSet", data, meta)
}

func resourceGeoMatchSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::GeoMatchSet", data, meta)
}