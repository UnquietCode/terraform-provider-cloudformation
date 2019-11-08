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

func ResourceWAFRegionalSizeConstraintSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceWAFRegionalSizeConstraintSetCreate,
		Read:   resourceWAFRegionalSizeConstraintSetRead,
		Update: resourceWAFRegionalSizeConstraintSetUpdate,
		Delete: resourceWAFRegionalSizeConstraintSetDelete,

		Schema: map[string]*schema.Schema{
			"size_constraints": {
				Type: schema.TypeList,
				Elem: propertySizeConstraint(),
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

func resourceWAFRegionalSizeConstraintSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::SizeConstraintSet", data, meta)
}

func resourceWAFRegionalSizeConstraintSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::SizeConstraintSet", data, meta)
}

func resourceWAFRegionalSizeConstraintSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::SizeConstraintSet", data, meta)
}

func resourceWAFRegionalSizeConstraintSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::SizeConstraintSet", data, meta)
}