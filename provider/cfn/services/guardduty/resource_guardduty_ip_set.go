// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGuardDutyIPSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceGuardDutyIPSetCreate,
		Read:   resourceGuardDutyIPSetRead,
		Update: resourceGuardDutyIPSetUpdate,
		Delete: resourceGuardDutyIPSetDelete,

		Schema: map[string]*schema.Schema{
			"format": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"activate": {
				Type: schema.TypeBool,
				Required: true,
			},
			"detector_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"location": {
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

func resourceGuardDutyIPSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::IPSet", ResourceGuardDutyIPSet(), data, meta)
}

func resourceGuardDutyIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::IPSet", ResourceGuardDutyIPSet(), data, meta)
}

func resourceGuardDutyIPSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::IPSet", ResourceGuardDutyIPSet(), data, meta)
}

func resourceGuardDutyIPSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::IPSet", data, meta)
}