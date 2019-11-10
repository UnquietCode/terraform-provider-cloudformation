// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEMRStep() *schema.Resource {
	return &schema.Resource{
		Create: resourceEMRStepCreate,
		Read:   resourceEMRStepRead,
		Update: resourceEMRStepUpdate,
		Delete: resourceEMRStepDelete,

		Schema: map[string]*schema.Schema{
			"action_on_failure": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"hadoop_jar_step": {
				Type: schema.TypeList,
				Elem: propertyStepHadoopJarStepConfig(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"job_flow_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEMRStepCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::Step", data, meta)
}

func resourceEMRStepRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::Step", data, meta)
}

func resourceEMRStepUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::Step", data, meta)
}

func resourceEMRStepDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::Step", data, meta)
}