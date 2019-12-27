// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttemplate.html

package inspector

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const inspectorAssessmentTemplateType string = "AWS::Inspector::AssessmentTemplate"

func DatasourceInspectorAssessmentTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceInspectorAssessmentTemplateRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceInspectorAssessmentTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(inspectorAssessmentTemplateType, DatasourceInspectorAssessmentTemplate(), data, meta)
}
