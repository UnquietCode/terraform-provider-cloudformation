// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceMemberCreate,
		Read:   resourceMemberRead,
		Update: resourceMemberUpdate,
		Delete: resourceMemberDelete,

		Schema: map[string]*schema.Schema{
			"status": {
				Type: schema.TypeString,
				Required: false,
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
				Required: false,
			},
			"disable_email_notification": {
				Type: schema.TypeBool,
				Required: false,
			},
			"detector_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMemberCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::Member", data, meta)
}

func resourceMemberRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::Member", data, meta)
}

func resourceMemberUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::Member", data, meta)
}

func resourceMemberDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::Member", data, meta)
}