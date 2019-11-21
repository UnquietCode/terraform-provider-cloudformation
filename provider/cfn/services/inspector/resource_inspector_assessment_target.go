// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html

package inspector

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const inspectorAssessmentTargetType string = "AWS::Inspector::AssessmentTarget"

func ResourceInspectorAssessmentTarget() *schema.Resource {
	return &schema.Resource{
		Read: resourceInspectorAssessmentTargetRead,
		Create: resourceInspectorAssessmentTargetCreate,
		Update: resourceInspectorAssessmentTargetUpdate,
		Delete: resourceInspectorAssessmentTargetDelete,
		CustomizeDiff: resourceInspectorAssessmentTargetCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"assessment_target_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_group_arn": {
				Type: schema.TypeString,
				Optional: true,
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

func resourceInspectorAssessmentTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(inspectorAssessmentTargetType, ResourceInspectorAssessmentTarget(), data, meta)
}

func resourceInspectorAssessmentTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(inspectorAssessmentTargetType, ResourceInspectorAssessmentTarget(), data, meta)
}

func resourceInspectorAssessmentTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(inspectorAssessmentTargetType, ResourceInspectorAssessmentTarget(), data, meta)
}

func resourceInspectorAssessmentTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(inspectorAssessmentTargetType, data, meta)
}

func resourceInspectorAssessmentTargetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(inspectorAssessmentTargetType, data, meta)
}
