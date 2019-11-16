// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eMRStepType string = "AWS::EMR::Step"

var eMRStepProperties map[string]string = map[string]string{
	"action_on_failure": "ActionOnFailure",
	"hadoop_jar_step": "HadoopJarStep",
	"job_flow_id": "JobFlowId",
	"name": "Name",
}

func ResourceEMRStep() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEMRStepExists,
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
				Type: schema.TypeList,
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
			},
		},
	}
}

func resourceEMRStepExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEMRStepRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eMRStepType, ResourceEMRStep(), data, meta)
}

func resourceEMRStepCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eMRStepType, ResourceEMRStep(), data, eMRStepProperties, meta)
}

func resourceEMRStepUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eMRStepType, ResourceEMRStep(), data, eMRStepProperties, meta)
}

func resourceEMRStepDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eMRStepType, data, meta)
}

func resourceEMRStepCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eMRStepType, data, meta)
}
