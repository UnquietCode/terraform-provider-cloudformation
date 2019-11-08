// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMMaintenanceWindow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSMMaintenanceWindowCreate,
		Read:   resourceSSMMaintenanceWindowRead,
		Update: resourceSSMMaintenanceWindowUpdate,
		Delete: resourceSSMMaintenanceWindowDelete,

		Schema: map[string]*schema.Schema{
			"start_date": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"allow_unassociated_targets": {
				Type: schema.TypeBool,
				Required: true,
			},
			"cutoff": {
				Type: schema.TypeInt,
				Required: true,
			},
			"schedule": {
				Type: schema.TypeString,
				Required: true,
			},
			"duration": {
				Type: schema.TypeInt,
				Required: true,
			},
			"end_date": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"schedule_timezone": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceSSMMaintenanceWindowCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::MaintenanceWindow", data, meta)
}

func resourceSSMMaintenanceWindowRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::MaintenanceWindow", data, meta)
}

func resourceSSMMaintenanceWindowUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::MaintenanceWindow", data, meta)
}

func resourceSSMMaintenanceWindowDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::MaintenanceWindow", data, meta)
}