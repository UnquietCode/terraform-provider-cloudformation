// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCampaignSchedule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"time_zone": {
				Type: schema.TypeString,
				Required: false,
			},
			"quiet_time": {
				Type: schema.TypeList,
				Elem: propertyCampaignQuietTime(),
				Required: false,
				MaxItems: 1,
			},
			"end_time": {
				Type: schema.TypeString,
				Required: false,
			},
			"start_time": {
				Type: schema.TypeString,
				Required: false,
			},
			"frequency": {
				Type: schema.TypeString,
				Required: false,
			},
			"event_filter": {
				Type: schema.TypeList,
				Elem: propertyCampaignCampaignEventFilter(),
				Required: false,
				MaxItems: 1,
			},
			"is_local_time": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}