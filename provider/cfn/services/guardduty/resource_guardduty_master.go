// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGuardDutyMaster() *schema.Resource {
	return &schema.Resource{
		Create: resourceGuardDutyMasterCreate,
		Read:   resourceGuardDutyMasterRead,
		Update: resourceGuardDutyMasterUpdate,
		Delete: resourceGuardDutyMasterDelete,

		Schema: map[string]*schema.Schema{
			"detector_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"master_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"invitation_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceGuardDutyMasterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::Master", data, meta)
}

func resourceGuardDutyMasterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::Master", data, meta)
}

func resourceGuardDutyMasterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::Master", data, meta)
}

func resourceGuardDutyMasterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::Master", data, meta)
}