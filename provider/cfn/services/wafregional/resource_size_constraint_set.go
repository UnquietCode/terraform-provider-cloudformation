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

func ResourceSizeConstraintSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSizeConstraintSetCreate,
		Read:   resourceSizeConstraintSetRead,
		Update: resourceSizeConstraintSetUpdate,
		Delete: resourceSizeConstraintSetDelete,

		Schema: map[string]*schema.Schema{
			"size_constraints": {
				Type: schema.TypeList,
				Elem: propertySizeConstraintSetSizeConstraint(),
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

func resourceSizeConstraintSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::SizeConstraintSet", data, meta)
}

func resourceSizeConstraintSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::SizeConstraintSet", data, meta)
}

func resourceSizeConstraintSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::SizeConstraintSet", data, meta)
}

func resourceSizeConstraintSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::SizeConstraintSet", data, meta)
}