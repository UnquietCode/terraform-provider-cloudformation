// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const guardDutyMemberType string = "AWS::GuardDuty::Member"

func ResourceGuardDutyMember() *schema.Resource {
	return &schema.Resource{
		Exists: resourceGuardDutyMemberExists,
		Read: resourceGuardDutyMemberRead,
		Create: resourceGuardDutyMemberCreate,
		Update: resourceGuardDutyMemberUpdate,
		Delete: resourceGuardDutyMemberDelete,
		CustomizeDiff: resourceGuardDutyMemberCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"status": {
				Type: schema.TypeString,
				Optional: true,
			},
			"member_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"email": {
				Type: schema.TypeString,
				Required: true,
			},
			"message": {
				Type: schema.TypeString,
				Optional: true,
			},
			"disable_email_notification": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"detector_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGuardDutyMemberExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceGuardDutyMemberRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(guardDutyMemberType, ResourceGuardDutyMember(), data, meta)
}

func resourceGuardDutyMemberCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(guardDutyMemberType, ResourceGuardDutyMember(), data, meta)
}

func resourceGuardDutyMemberUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(guardDutyMemberType, ResourceGuardDutyMember(), data, meta)
}

func resourceGuardDutyMemberDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(guardDutyMemberType, data, meta)
}

func resourceGuardDutyMemberCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(guardDutyMemberType, data, meta)
}
