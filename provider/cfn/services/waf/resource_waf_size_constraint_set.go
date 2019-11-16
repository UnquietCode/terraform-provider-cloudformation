// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const wAFSizeConstraintSetType string = "AWS::WAF::SizeConstraintSet"

var wAFSizeConstraintSetProperties map[string]string = map[string]string{
	"name": "Name",
	"size_constraints": "SizeConstraints",
}

func ResourceWAFSizeConstraintSet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFSizeConstraintSetExists,
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
			},
		},
	}
}

func resourceWAFSizeConstraintSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceWAFSizeConstraintSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(wAFSizeConstraintSetType, ResourceWAFSizeConstraintSet(), data, meta)
}

func resourceWAFSizeConstraintSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(wAFSizeConstraintSetType, ResourceWAFSizeConstraintSet(), data, wAFSizeConstraintSetProperties, meta)
}

func resourceWAFSizeConstraintSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(wAFSizeConstraintSetType, ResourceWAFSizeConstraintSet(), data, wAFSizeConstraintSetProperties, meta)
}

func resourceWAFSizeConstraintSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(wAFSizeConstraintSetType, data, meta)
}

func resourceWAFSizeConstraintSetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(wAFSizeConstraintSetType, data, meta)
}
