// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html

package waf

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFSizeConstraintSetType string = "AWS::WAF::SizeConstraintSet"

func ResourceWAFSizeConstraintSet() *schema.Resource {
	return &schema.Resource{
		Read: resourceWAFSizeConstraintSetRead,
		Create: resourceWAFSizeConstraintSetCreate,
		Update: resourceWAFSizeConstraintSetUpdate,
		Delete: resourceWAFSizeConstraintSetDelete,
		CustomizeDiff: resourceWAFSizeConstraintSetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"size_constraints": {
				Type: schema.TypeSet,
				Elem: propertySizeConstraintSetSizeConstraint(),
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceWAFSizeConstraintSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFSizeConstraintSetType, ResourceWAFSizeConstraintSet(), data, meta)
}

func resourceWAFSizeConstraintSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFSizeConstraintSetType, ResourceWAFSizeConstraintSet(), data, meta)
}

func resourceWAFSizeConstraintSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFSizeConstraintSetType, ResourceWAFSizeConstraintSet(), data, meta)
}

func resourceWAFSizeConstraintSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFSizeConstraintSetType, data, meta)
}

func resourceWAFSizeConstraintSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFSizeConstraintSetType, data, meta)
}
