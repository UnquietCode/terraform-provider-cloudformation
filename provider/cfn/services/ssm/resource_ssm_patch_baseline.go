// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

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
				Required: false,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"approval_rules": {
				Type: schema.TypeList,
				Elem: propertyRuleGroup(),
				Required: false,
				MaxItems: 1,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertyPatchSource(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"rejected_patches": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"approved_patches": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"rejected_patches_action": {
				Type: schema.TypeString,
				Required: false,
			},
			"patch_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"approved_patches_compliance_level": {
				Type: schema.TypeString,
				Required: false,
			},
			"approved_patches_enable_non_security": {
				Type: schema.TypeBool,
				Required: false,
			},
			"global_filters": {
				Type: schema.TypeList,
				Elem: propertyPatchFilterGroup(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
		},
	}
}

func resourceSSMPatchBaselineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::PatchBaseline", data, meta)
}

func resourceSSMPatchBaselineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::PatchBaseline", data, meta)
}

func resourceSSMPatchBaselineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::PatchBaseline", data, meta)
}

func resourceSSMPatchBaselineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::PatchBaseline", data, meta)
}