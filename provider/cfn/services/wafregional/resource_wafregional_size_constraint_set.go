// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFRegionalSizeConstraintSet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFRegionalSizeConstraintSetExists,
		Read:   resourceWAFRegionalSizeConstraintSetRead,
		Create: resourceWAFRegionalSizeConstraintSetCreate,
		Update: resourceWAFRegionalSizeConstraintSetUpdate,
		Delete: resourceWAFRegionalSizeConstraintSetDelete,
		CustomizeDiff: resourceWAFRegionalSizeConstraintSetCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"size_constraints": {
				Type: schema.TypeList,
				Elem: propertySizeConstraintSetSizeConstraint(),
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

func resourceWAFRegionalSizeConstraintSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFRegionalSizeConstraintSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::WAFRegional::SizeConstraintSet", ResourceWAFRegionalSizeConstraintSet(), data, meta)
}

func resourceWAFRegionalSizeConstraintSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAFRegional::SizeConstraintSet", ResourceWAFRegionalSizeConstraintSet(), data, meta)
}

func resourceWAFRegionalSizeConstraintSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAFRegional::SizeConstraintSet", ResourceWAFRegionalSizeConstraintSet(), data, meta)
}

func resourceWAFRegionalSizeConstraintSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAFRegional::SizeConstraintSet", data, meta)
}

func resourceWAFRegionalSizeConstraintSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}