// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStep() *schema.Resource {
	return &schema.Resource{
		Create: resourceStepCreate,
		Read:   resourceStepRead,
		Update: resourceStepUpdate,
		Delete: resourceStepDelete,

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

func resourceStepCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EMR::Step", data, meta)
}

func resourceStepRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EMR::Step", data, meta)
}

func resourceStepUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EMR::Step", data, meta)
}

func resourceStepDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EMR::Step", data, meta)
}