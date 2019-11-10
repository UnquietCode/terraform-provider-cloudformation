// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatintelset.html

package guardduty

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGuardDutyThreatIntelSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceGuardDutyThreatIntelSetCreate,
		Read:   resourceGuardDutyThreatIntelSetRead,
		Update: resourceGuardDutyThreatIntelSetUpdate,
		Delete: resourceGuardDutyThreatIntelSetDelete,

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
		},
	}
}

func resourceGuardDutyThreatIntelSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::ThreatIntelSet", data, meta)
}

func resourceGuardDutyThreatIntelSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::ThreatIntelSet", data, meta)
}

func resourceGuardDutyThreatIntelSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::ThreatIntelSet", data, meta)
}

func resourceGuardDutyThreatIntelSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::ThreatIntelSet", data, meta)
}