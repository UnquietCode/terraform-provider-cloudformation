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

func ResourceIPSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPSetCreate,
		Read:   resourceIPSetRead,
		Update: resourceIPSetUpdate,
		Delete: resourceIPSetDelete,

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

func resourceIPSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::IPSet", data, meta)
}

func resourceIPSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::IPSet", data, meta)
}

func resourceIPSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::IPSet", data, meta)
}

func resourceIPSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::IPSet", data, meta)
}