// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
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
	return plugin.ResourceRead("AWS::SSM::MaintenanceWindow", ResourceSSMMaintenanceWindow(), data, meta)
}

func resourceSSMMaintenanceWindowCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::MaintenanceWindow", ResourceSSMMaintenanceWindow(), data, meta)
}

func resourceSSMMaintenanceWindowUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::MaintenanceWindow", ResourceSSMMaintenanceWindow(), data, meta)
}

func resourceSSMMaintenanceWindowDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::MaintenanceWindow", data, meta)
}

func resourceSSMMaintenanceWindowCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
