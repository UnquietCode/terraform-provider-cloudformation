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

func ResourceSSMMaintenanceWindowTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSMMaintenanceWindowTargetCreate,
		Read:   resourceSSMMaintenanceWindowTargetRead,
		Update: resourceSSMMaintenanceWindowTargetUpdate,
		Delete: resourceSSMMaintenanceWindowTargetDelete,

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
				Elem: propertyTargets(),
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceSSMMaintenanceWindowTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::MaintenanceWindowTarget", data, meta)
}

func resourceSSMMaintenanceWindowTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::MaintenanceWindowTarget", data, meta)
}

func resourceSSMMaintenanceWindowTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::MaintenanceWindowTarget", data, meta)
}

func resourceSSMMaintenanceWindowTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::MaintenanceWindowTarget", data, meta)
}