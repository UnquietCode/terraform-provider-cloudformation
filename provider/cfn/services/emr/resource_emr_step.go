// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html

package emr

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eMRStepType string = "AWS::EMR::Step"

func ResourceEMRStep() *schema.Resource {
	return &schema.Resource{
		Read: resourceEMRStepRead,
		Create: resourceEMRStepCreate,
		Update: resourceEMRStepUpdate,
		Delete: resourceEMRStepDelete,
		CustomizeDiff: resourceEMRStepCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"action_on_failure": {
				Type: schema.TypeString,
				Required: true,
			},
			"hadoop_jar_step": {
				Type: schema.TypeSet,
				Elem: propertyStepHadoopJarStepConfig(),
				Required: true,
				MaxItems: 1,
			},
			"job_flow_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
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

func resourceEMRStepRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRStepType, ResourceEMRStep(), data, meta)
}

func resourceEMRStepCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eMRStepType, ResourceEMRStep(), data, meta)
}

func resourceEMRStepUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eMRStepType, ResourceEMRStep(), data, meta)
}

func resourceEMRStepDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eMRStepType, data, meta)
}

func resourceEMRStepCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eMRStepType, data, meta)
}
