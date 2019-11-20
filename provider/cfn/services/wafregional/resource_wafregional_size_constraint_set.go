// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-sizeconstraintset.html

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFRegionalSizeConstraintSetType string = "AWS::WAFRegional::SizeConstraintSet"

func ResourceWAFRegionalSizeConstraintSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFRegionalSizeConstraintSetRead,
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

func resourceWAFRegionalSizeConstraintSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFRegionalSizeConstraintSetType, ResourceWAFRegionalSizeConstraintSet(), data, meta)
}

func resourceWAFRegionalSizeConstraintSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFRegionalSizeConstraintSetType, ResourceWAFRegionalSizeConstraintSet(), data, meta)
}

func resourceWAFRegionalSizeConstraintSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFRegionalSizeConstraintSetType, ResourceWAFRegionalSizeConstraintSet(), data, meta)
}

func resourceWAFRegionalSizeConstraintSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFRegionalSizeConstraintSetType, data, meta)
}

func resourceWAFRegionalSizeConstraintSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFRegionalSizeConstraintSetType, data, meta)
}
