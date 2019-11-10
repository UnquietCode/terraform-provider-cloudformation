// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMaintenanceWindowTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceMaintenanceWindowTargetCreate,
		Read:   resourceMaintenanceWindowTargetRead,
		Update: resourceMaintenanceWindowTargetUpdate,
		Delete: resourceMaintenanceWindowTargetDelete,

		Schema: map[string]*schema.Schema{
			"owner_information": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"window_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"targets": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindowTargetTargets(),
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceMaintenanceWindowTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::MaintenanceWindowTarget", data, meta)
}

func resourceMaintenanceWindowTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::MaintenanceWindowTarget", data, meta)
}

func resourceMaintenanceWindowTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::MaintenanceWindowTarget", data, meta)
}

func resourceMaintenanceWindowTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::MaintenanceWindowTarget", data, meta)
}