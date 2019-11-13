// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html

package waf

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWAFSizeConstraintSet() *schema.Resource {
	return &schema.Resource{
		Exists: resourceWAFSizeConstraintSetExists,
		Read:   resourceWAFSizeConstraintSetRead,
		Create: resourceWAFSizeConstraintSetCreate,
		Update: resourceWAFSizeConstraintSetUpdate,
		Delete: resourceWAFSizeConstraintSetDelete,
		
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
	return plugin.ResourceRead("AWS::WAF::SizeConstraintSet", ResourceWAFSizeConstraintSet(), data, meta)
}

func resourceWAFSizeConstraintSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::WAF::SizeConstraintSet", ResourceWAFSizeConstraintSet(), data, meta)
}

func resourceWAFSizeConstraintSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::WAF::SizeConstraintSet", ResourceWAFSizeConstraintSet(), data, meta)
}

func resourceWAFSizeConstraintSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::WAF::SizeConstraintSet", data, meta)
}