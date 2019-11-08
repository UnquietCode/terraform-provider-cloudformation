// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMMaintenanceWindowTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSMMaintenanceWindowTaskCreate,
		Read:   resourceSSMMaintenanceWindowTaskRead,
		Update: resourceSSMMaintenanceWindowTaskUpdate,
		Delete: resourceSSMMaintenanceWindowTaskDelete,

		Schema: map[string]*schema.Schema{
			"max_errors": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"service_role_arn": {
				Type: schema.TypeString,
				Required: false,
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
				Elem: propertyTarget(),
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
			"task_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"task_invocation_parameters": {
				Type: schema.TypeList,
				Elem: propertyTaskInvocationParameters(),
				Required: false,
				MaxItems: 1,
			},
			"window_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"task_parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"task_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"logging_info": {
				Type: schema.TypeList,
				Elem: propertyLoggingInfo(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceSSMMaintenanceWindowTaskCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::MaintenanceWindowTask", data, meta)
}

func resourceSSMMaintenanceWindowTaskRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::MaintenanceWindowTask", data, meta)
}

func resourceSSMMaintenanceWindowTaskUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::MaintenanceWindowTask", data, meta)
}

func resourceSSMMaintenanceWindowTaskDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::MaintenanceWindowTask", data, meta)
}