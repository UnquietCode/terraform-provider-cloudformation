// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGuardDutyMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceGuardDutyMemberCreate,
		Read:   resourceGuardDutyMemberRead,
		Update: resourceGuardDutyMemberUpdate,
		Delete: resourceGuardDutyMemberDelete,

		Schema: map[string]*schema.Schema{
			"status": {
				Type: schema.TypeString,
				Optional: true,
			},
			"member_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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
				ForceNew: true,
			},
		},
	}
}

func resourceGuardDutyMemberCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::Member", data, meta)
}

func resourceGuardDutyMemberRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::Member", data, meta)
}

func resourceGuardDutyMemberUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::Member", data, meta)
}

func resourceGuardDutyMemberDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::Member", data, meta)
}