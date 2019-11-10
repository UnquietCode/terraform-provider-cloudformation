// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMPatchBaseline() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSMPatchBaselineCreate,
		Read:   resourceSSMPatchBaselineRead,
		Update: resourceSSMPatchBaselineUpdate,
		Delete: resourceSSMPatchBaselineDelete,

		Schema: map[string]*schema.Schema{
			"operating_system": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"approval_rules": {
				Type: schema.TypeList,
				Elem: propertyPatchBaselineRuleGroup(),
				Optional: true,
				MaxItems: 1,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertyPatchBaselinePatchSource(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"rejected_patches": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"approved_patches": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"rejected_patches_action": {
				Type: schema.TypeString,
				Optional: true,
			},
			"patch_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"approved_patches_compliance_level": {
				Type: schema.TypeString,
				Optional: true,
			},
			"approved_patches_enable_non_security": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"global_filters": {
				Type: schema.TypeList,
				Elem: propertyPatchBaselinePatchFilterGroup(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
		},
	}
}

func resourceSSMPatchBaselineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::PatchBaseline", ResourceSSMPatchBaseline(), data, meta)
}

func resourceSSMPatchBaselineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::PatchBaseline", ResourceSSMPatchBaseline(), data, meta)
}

func resourceSSMPatchBaselineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::PatchBaseline", ResourceSSMPatchBaseline(), data, meta)
}

func resourceSSMPatchBaselineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::PatchBaseline", data, meta)
}