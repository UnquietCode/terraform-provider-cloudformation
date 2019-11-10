// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.html

package ssm

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMaintenanceWindowTaskTaskInvocationParameters(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"maintenance_window_run_command_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTaskMaintenanceWindowRunCommandParameters(),
				Required: false,
				MaxItems: 1,
			},
			"maintenance_window_automation_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTaskMaintenanceWindowAutomationParameters(),
				Required: false,
				MaxItems: 1,
			},
			"maintenance_window_step_functions_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTaskMaintenanceWindowStepFunctionsParameters(),
				Required: false,
				MaxItems: 1,
			},
			"maintenance_window_lambda_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTaskMaintenanceWindowLambdaParameters(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}