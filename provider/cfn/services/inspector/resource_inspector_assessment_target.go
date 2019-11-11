// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-inspector-assessmenttarget.html

package inspector

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceInspectorAssessmentTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceInspectorAssessmentTargetCreate,
		Read:   resourceInspectorAssessmentTargetRead,
		Update: resourceInspectorAssessmentTargetUpdate,
		Delete: resourceInspectorAssessmentTargetDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"assessment_target_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"resource_group_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceInspectorAssessmentTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Inspector::AssessmentTarget", ResourceInspectorAssessmentTarget(), data, meta)
}

func resourceInspectorAssessmentTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Inspector::AssessmentTarget", ResourceInspectorAssessmentTarget(), data, meta)
}

func resourceInspectorAssessmentTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Inspector::AssessmentTarget", ResourceInspectorAssessmentTarget(), data, meta)
}

func resourceInspectorAssessmentTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Inspector::AssessmentTarget", data, meta)
}