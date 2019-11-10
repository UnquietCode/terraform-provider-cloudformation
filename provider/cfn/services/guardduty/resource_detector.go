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

func ResourceDetector() *schema.Resource {
	return &schema.Resource{
		Create: resourceDetectorCreate,
		Read:   resourceDetectorRead,
		Update: resourceDetectorUpdate,
		Delete: resourceDetectorDelete,

		Schema: map[string]*schema.Schema{
			"finding_publishing_frequency": {
				Type: schema.TypeString,
				Required: false,
			},
			"enable": {
				Type: schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceDetectorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::GuardDuty::Detector", data, meta)
}

func resourceDetectorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::GuardDuty::Detector", data, meta)
}

func resourceDetectorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::GuardDuty::Detector", data, meta)
}

func resourceDetectorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::GuardDuty::Detector", data, meta)
}