// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html

package inspector

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceInspectorAssessmentTemplate() *schema.Resource {
	return &schema.Resource{
		Exists: resourceInspectorAssessmentTemplateExists,
		Read: resourceInspectorAssessmentTemplateRead,
		Create: resourceInspectorAssessmentTemplateCreate,
		Update: resourceInspectorAssessmentTemplateUpdate,
		Delete: resourceInspectorAssessmentTemplateDelete,
		CustomizeDiff: resourceInspectorAssessmentTemplateCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"assessment_target_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"duration_in_seconds": {
				Type: schema.TypeInt,
				Required: true,
			},
			"assessment_template_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rules_package_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"user_attributes_for_findings": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceInspectorAssessmentTemplateExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceInspectorAssessmentTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Inspector::AssessmentTemplate", ResourceInspectorAssessmentTemplate(), data, meta)
}

func resourceInspectorAssessmentTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Inspector::AssessmentTemplate", ResourceInspectorAssessmentTemplate(), data, meta)
}

func resourceInspectorAssessmentTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Inspector::AssessmentTemplate", ResourceInspectorAssessmentTemplate(), data, meta)
}

func resourceInspectorAssessmentTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Inspector::AssessmentTemplate", data, meta)
}

func resourceInspectorAssessmentTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::Inspector::AssessmentTemplate", data, meta)
}

