// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-statemachine.html

package stepfunctions

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const stepFunctionsStateMachineType string = "AWS::StepFunctions::StateMachine"

func DatasourceStepFunctionsStateMachine() *schema.Resource {
	return &schema.Resource{
		Read: datasourceStepFunctionsStateMachineRead,
		
		Schema: map[string]*schema.Schema{
			"definition_string": {
				Type: schema.TypeString,
				Required: true,
			},
			"logging_configuration": {
				Type: schema.TypeList,
				Elem: propertyStateMachineLoggingConfiguration(),
				Optional: true,
				MaxItems: 1,
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
			"state_machine_type": {
				Type: schema.TypeString,
				Optional: true,
			},
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

func datasourceStepFunctionsStateMachineRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(stepFunctionsStateMachineType, DatasourceStepFunctionsStateMachine(), data, meta)
}
