// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindow.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sSMMaintenanceWindowType string = "AWS::SSM::MaintenanceWindow"

var sSMMaintenanceWindowProperties map[string]string = map[string]string{
	"start_date": "StartDate",
	"description": "Description",
	"allow_unassociated_targets": "AllowUnassociatedTargets",
	"cutoff": "Cutoff",
	"schedule": "Schedule",
	"duration": "Duration",
	"end_date": "EndDate",
	"tags": "Tags",
	"name": "Name",
	"schedule_timezone": "ScheduleTimezone",
}

func ResourceSSMMaintenanceWindow() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSSMMaintenanceWindowExists,
		Read: resourceSSMMaintenanceWindowRead,
		Create: resourceSSMMaintenanceWindowCreate,
		Update: resourceSSMMaintenanceWindowUpdate,
		Delete: resourceSSMMaintenanceWindowDelete,
		CustomizeDiff: resourceSSMMaintenanceWindowCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"start_date": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
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
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"schedule_timezone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSSMMaintenanceWindowExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSSMMaintenanceWindowRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sSMMaintenanceWindowType, ResourceSSMMaintenanceWindow(), data, meta)
}

func resourceSSMMaintenanceWindowCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sSMMaintenanceWindowType, ResourceSSMMaintenanceWindow(), data, sSMMaintenanceWindowProperties, meta)
}

func resourceSSMMaintenanceWindowUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sSMMaintenanceWindowType, ResourceSSMMaintenanceWindow(), data, sSMMaintenanceWindowProperties, meta)
}

func resourceSSMMaintenanceWindowDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sSMMaintenanceWindowType, data, meta)
}

func resourceSSMMaintenanceWindowCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sSMMaintenanceWindowType, data, meta)
}
