// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCampaignWriteTreatmentResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"treatment_description": {
				Type: schema.TypeString,
				Required: false,
			},
			"message_configuration": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessageConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"schedule": {
				Type: schema.TypeList,
				Elem: propertyCampaignSchedule(),
				Required: false,
				MaxItems: 1,
			},
			"size_percent": {
				Type: schema.TypeInt,
				Required: false,
			},
			"treatment_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}