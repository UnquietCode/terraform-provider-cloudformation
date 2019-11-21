// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html

package ssm

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sSMPatchBaselineType string = "AWS::SSM::PatchBaseline"

func ResourceSSMPatchBaseline() *schema.Resource {
	return &schema.Resource{
		Read: resourceSSMPatchBaselineRead,
		Create: resourceSSMPatchBaselineCreate,
		Update: resourceSSMPatchBaselineUpdate,
		Delete: resourceSSMPatchBaselineDelete,
		CustomizeDiff: resourceSSMPatchBaselineCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"operating_system": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"approval_rules": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
				Elem: propertyPatchBaselinePatchFilterGroup(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceSSMPatchBaselineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sSMPatchBaselineType, ResourceSSMPatchBaseline(), data, meta)
}

func resourceSSMPatchBaselineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sSMPatchBaselineType, ResourceSSMPatchBaseline(), data, meta)
}

func resourceSSMPatchBaselineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sSMPatchBaselineType, ResourceSSMPatchBaseline(), data, meta)
}

func resourceSSMPatchBaselineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sSMPatchBaselineType, data, meta)
}

func resourceSSMPatchBaselineCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sSMPatchBaselineType, data, meta)
}
