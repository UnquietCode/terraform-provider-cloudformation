// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sSMMaintenanceWindowTaskType string = "AWS::SSM::MaintenanceWindowTask"

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
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
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
	return plugin.ResourceRead(sSMMaintenanceWindowTaskType, ResourceSSMMaintenanceWindowTask(), data, meta)
}

func resourceSSMMaintenanceWindowTaskCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sSMMaintenanceWindowTaskType, ResourceSSMMaintenanceWindowTask(), data, meta)
}

func resourceSSMMaintenanceWindowTaskUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sSMMaintenanceWindowTaskType, ResourceSSMMaintenanceWindowTask(), data, meta)
}

func resourceSSMMaintenanceWindowTaskDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sSMMaintenanceWindowTaskType, data, meta)
}

func resourceSSMMaintenanceWindowTaskCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sSMMaintenanceWindowTaskType, data, meta)
}
