// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const sSMPatchBaselineType string = "AWS::SSM::PatchBaseline"

var sSMPatchBaselineProperties map[string]string = map[string]string{
	"operating_system": "OperatingSystem",
	"description": "Description",
	"approval_rules": "ApprovalRules",
	"sources": "Sources",
	"name": "Name",
	"rejected_patches": "RejectedPatches",
	"approved_patches": "ApprovedPatches",
	"rejected_patches_action": "RejectedPatchesAction",
	"patch_groups": "PatchGroups",
	"approved_patches_compliance_level": "ApprovedPatchesComplianceLevel",
	"approved_patches_enable_non_security": "ApprovedPatchesEnableNonSecurity",
	"global_filters": "GlobalFilters",
	"tags": "Tags",
}

func ResourceSSMPatchBaseline() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSSMPatchBaselineExists,
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
			},
		},
	}
}

func resourceSSMPatchBaselineExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSSMPatchBaselineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sSMPatchBaselineType, ResourceSSMPatchBaseline(), data, meta)
}

func resourceSSMPatchBaselineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sSMPatchBaselineType, ResourceSSMPatchBaseline(), data, sSMPatchBaselineProperties, meta)
}

func resourceSSMPatchBaselineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sSMPatchBaselineType, ResourceSSMPatchBaseline(), data, sSMPatchBaselineProperties, meta)
}

func resourceSSMPatchBaselineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sSMPatchBaselineType, data, meta)
}

func resourceSSMPatchBaselineCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sSMPatchBaselineType, data, meta)
}
