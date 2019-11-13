// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtarget.html

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSSMMaintenanceWindowTarget() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSSMMaintenanceWindowTargetExists,
		Read:   resourceSSMMaintenanceWindowTargetRead,
		Create: resourceSSMMaintenanceWindowTargetCreate,
		Update: resourceSSMMaintenanceWindowTargetUpdate,
		Delete: resourceSSMMaintenanceWindowTargetDelete,
		CustomizeDiff: resourceSSMMaintenanceWindowTargetCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"owner_information": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"window_id": {
				Type: schema.TypeString,
				Required: true,
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

func resourceSSMMaintenanceWindowTargetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSSMMaintenanceWindowTargetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SSM::MaintenanceWindowTarget", ResourceSSMMaintenanceWindowTarget(), data, meta)
}

func resourceSSMMaintenanceWindowTargetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SSM::MaintenanceWindowTarget", ResourceSSMMaintenanceWindowTarget(), data, meta)
}

func resourceSSMMaintenanceWindowTargetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SSM::MaintenanceWindowTarget", ResourceSSMMaintenanceWindowTarget(), data, meta)
}

func resourceSSMMaintenanceWindowTargetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SSM::MaintenanceWindowTarget", data, meta)
}

func resourceSSMMaintenanceWindowTargetCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}