// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const inspectorAssessmentTemplateType string = "AWS::Inspector::AssessmentTemplate"

func ResourceInspectorAssessmentTemplate() *schema.Resource {
	return &schema.Resource{
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
			"user_attributes_for_findings": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceInspectorAssessmentTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(inspectorAssessmentTemplateType, ResourceInspectorAssessmentTemplate(), data, meta)
}

func resourceInspectorAssessmentTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(inspectorAssessmentTemplateType, ResourceInspectorAssessmentTemplate(), data, meta)
}

func resourceInspectorAssessmentTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(inspectorAssessmentTemplateType, ResourceInspectorAssessmentTemplate(), data, meta)
}

func resourceInspectorAssessmentTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(inspectorAssessmentTemplateType, data, meta)
}

func resourceInspectorAssessmentTemplateCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(inspectorAssessmentTemplateType, data, meta)
}
