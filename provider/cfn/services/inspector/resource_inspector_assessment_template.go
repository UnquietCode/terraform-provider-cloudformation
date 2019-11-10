// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package inspector

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceInspectorAssessmentTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceInspectorAssessmentTemplateCreate,
		Read:   resourceInspectorAssessmentTemplateRead,
		Update: resourceInspectorAssessmentTemplateUpdate,
		Delete: resourceInspectorAssessmentTemplateDelete,

		Schema: map[string]*schema.Schema{
			"assessment_target_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"duration_in_seconds": {
				Type: schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"assessment_template_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"rules_package_arns": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"user_attributes_for_findings": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceInspectorAssessmentTemplateCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Inspector::AssessmentTemplate", data, meta)
}

func resourceInspectorAssessmentTemplateRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Inspector::AssessmentTemplate", data, meta)
}

func resourceInspectorAssessmentTemplateUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Inspector::AssessmentTemplate", data, meta)
}

func resourceInspectorAssessmentTemplateDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Inspector::AssessmentTemplate", data, meta)
}