// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html

package pinpoint

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCampaignSchedule(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"time_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"quiet_time": {
				Type: schema.TypeList,
				Elem: propertyCampaignQuietTime(),
				Optional: true,
				MaxItems: 1,
			},
			"end_time": {
				Type: schema.TypeString,
				Optional: true,
			},
			"start_time": {
				Type: schema.TypeString,
				Optional: true,
			},
			"frequency": {
				Type: schema.TypeString,
				Optional: true,
			},
			"event_filter": {
				Type: schema.TypeList,
				Elem: propertyCampaignCampaignEventFilter(),
				Optional: true,
				MaxItems: 1,
			},
			"is_local_time": {
				Type: schema.TypeBool,
				Optional: true,
			},
		},
	}
}
