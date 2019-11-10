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

func ResourceThreatIntelSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceThreatIntelSetCreate,
		Read:   resourceThreatIntelSetRead,
		Update: resourceThreatIntelSetUpdate,
		Delete: resourceThreatIntelSetDelete,

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
				Required: false,
			},
			"location": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceThreatIntelSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::ThreatIntelSet", data, meta)
}

func resourceThreatIntelSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::ThreatIntelSet", data, meta)
}

func resourceThreatIntelSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::ThreatIntelSet", data, meta)
}

func resourceThreatIntelSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::ThreatIntelSet", data, meta)
}