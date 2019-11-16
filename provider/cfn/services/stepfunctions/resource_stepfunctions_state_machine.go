// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html

package stepfunctions

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const stepFunctionsStateMachineType string = "AWS::StepFunctions::StateMachine"

var stepFunctionsStateMachineProperties map[string]string = map[string]string{
	"definition_string": "DefinitionString",
	"state_machine_name": "StateMachineName",
	"role_arn": "RoleArn",
	"tags": "Tags",
}

func ResourceStepFunctionsStateMachine() *schema.Resource {
	return &schema.Resource{
		Exists: resourceStepFunctionsStateMachineExists,
		Read: resourceStepFunctionsStateMachineRead,
		Create: resourceStepFunctionsStateMachineCreate,
		Update: resourceStepFunctionsStateMachineUpdate,
		Delete: resourceStepFunctionsStateMachineDelete,
		CustomizeDiff: resourceStepFunctionsStateMachineCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"definition_string": {
				Type: schema.TypeString,
				Required: true,
			},
			"state_machine_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyStateMachineTagsEntry(),
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

func resourceStepFunctionsStateMachineExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceStepFunctionsStateMachineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(stepFunctionsStateMachineType, ResourceStepFunctionsStateMachine(), data, meta)
}

func resourceStepFunctionsStateMachineCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(stepFunctionsStateMachineType, ResourceStepFunctionsStateMachine(), data, stepFunctionsStateMachineProperties, meta)
}

func resourceStepFunctionsStateMachineUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(stepFunctionsStateMachineType, ResourceStepFunctionsStateMachine(), data, stepFunctionsStateMachineProperties, meta)
}

func resourceStepFunctionsStateMachineDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(stepFunctionsStateMachineType, data, meta)
}

func resourceStepFunctionsStateMachineCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(stepFunctionsStateMachineType, data, meta)
}
