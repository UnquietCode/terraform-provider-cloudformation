// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTaskInvocationParameters() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_window_run_command_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowRunCommandParameters(),
				Required: false,
				MaxItems: 1,
			},
			"maintenance_window_automation_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowAutomationParameters(),
				Required: false,
				MaxItems: 1,
			},
			"maintenance_window_step_functions_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowStepFunctionsParameters(),
				Required: false,
				MaxItems: 1,
			},
			"maintenance_window_lambda_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowLambdaParameters(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}