// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMMaintenanceWindowTask() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSSMMaintenanceWindowTaskExists,
		Read: resourceSSMMaintenanceWindowTaskRead,
		Create: resourceSSMMaintenanceWindowTaskCreate,
		Update: resourceSSMMaintenanceWindowTaskUpdate,
		Delete: resourceSSMMaintenanceWindowTaskDelete,
		CustomizeDiff: resourceSSMMaintenanceWindowTaskCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"max_errors": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type: schema.TypeInt,
				Required: true,
			},
			"max_concurrency": {
				Type: schema.TypeString,
				Required: true,
			},
			"targets": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTaskTarget(),
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"task_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"task_invocation_parameters": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTaskTaskInvocationParameters(),
				Optional: true,
				MaxItems: 1,
			},
			"window_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"task_parameters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"task_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"logging_info": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTaskLoggingInfo(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSSMMaintenanceWindowTaskExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSSMMaintenanceWindowTaskRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::MaintenanceWindowTask", ResourceSSMMaintenanceWindowTask(), data, meta)
}

func resourceSSMMaintenanceWindowTaskCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::MaintenanceWindowTask", ResourceSSMMaintenanceWindowTask(), data, meta)
}

func resourceSSMMaintenanceWindowTaskUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::MaintenanceWindowTask", ResourceSSMMaintenanceWindowTask(), data, meta)
}

func resourceSSMMaintenanceWindowTaskDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::MaintenanceWindowTask", data, meta)
}

func resourceSSMMaintenanceWindowTaskCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

