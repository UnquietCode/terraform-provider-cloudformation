// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html

package ssm

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sSMMaintenanceWindowTaskType string = "AWS::SSM::MaintenanceWindowTask"

func DatasourceSSMMaintenanceWindowTask() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSSMMaintenanceWindowTaskRead,
		
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
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceSSMMaintenanceWindowTaskRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sSMMaintenanceWindowTaskType, DatasourceSSMMaintenanceWindowTask(), data, meta)
}
