// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const guardDutyMasterType string = "AWS::GuardDuty::Master"

func ResourceGuardDutyMaster() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGuardDutyMasterExists,
		Read: resourceGuardDutyMasterRead,
		Create: resourceGuardDutyMasterCreate,
		Update: resourceGuardDutyMasterUpdate,
		Delete: resourceGuardDutyMasterDelete,
		CustomizeDiff: resourceGuardDutyMasterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"detector_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"master_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"invitation_id": {
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

func resourceGuardDutyMasterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGuardDutyMasterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(guardDutyMasterType, ResourceGuardDutyMaster(), data, meta)
}

func resourceGuardDutyMasterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(guardDutyMasterType, ResourceGuardDutyMaster(), data, meta)
}

func resourceGuardDutyMasterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(guardDutyMasterType, ResourceGuardDutyMaster(), data, meta)
}

func resourceGuardDutyMasterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(guardDutyMasterType, data, meta)
}

func resourceGuardDutyMasterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(guardDutyMasterType, data, meta)
}
