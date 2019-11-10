// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package stepfunctions

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStepFunctionsStateMachine() *schema.Resource {
	return &schema.Resource{
		Create: resourceStepFunctionsStateMachineCreate,
		Read:   resourceStepFunctionsStateMachineRead,
		Update: resourceStepFunctionsStateMachineUpdate,
		Delete: resourceStepFunctionsStateMachineDelete,

		Schema: map[string]*schema.Schema{
			"definition_string": {
				Type: schema.TypeString,
				Required: true,
			},
			"state_machine_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyStateMachineTagsEntry(),
				Required: false,
			},
		},
	}
}

func resourceStepFunctionsStateMachineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::StepFunctions::StateMachine", data, meta)
}

func resourceStepFunctionsStateMachineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::StepFunctions::StateMachine", data, meta)
}

func resourceStepFunctionsStateMachineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::StepFunctions::StateMachine", data, meta)
}

func resourceStepFunctionsStateMachineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::StepFunctions::StateMachine", data, meta)
}