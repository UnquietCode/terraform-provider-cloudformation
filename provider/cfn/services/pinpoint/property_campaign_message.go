// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-message.html

package pinpoint

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCampaignMessage(extras...string) *schema.Resource {
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
			"json_body": {
				Type: schema.TypeString,
				Optional: true,
			},
			"action": {
				Type: schema.TypeString,
				Optional: true,
			},
			"media_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"time_to_live": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"image_small_icon_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"title": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image_icon_url": {
				Type: schema.TypeString,
				Optional: true,
			},
			"silent_push": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"body": {
				Type: schema.TypeString,
				Optional: true,
			},
			"raw_content": {
				Type: schema.TypeString,
				Optional: true,
			},
			"url": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
