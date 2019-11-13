// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-messageconfiguration.html

package pinpoint

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCampaignMessageConfiguration(extras...string) *schema.Resource {
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
			"apns_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Optional: true,
				MaxItems: 1,
			},
			"baidu_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Optional: true,
				MaxItems: 1,
			},
			"default_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Optional: true,
				MaxItems: 1,
			},
			"email_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignCampaignEmailMessage(),
				Optional: true,
				MaxItems: 1,
			},
			"gcm_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Optional: true,
				MaxItems: 1,
			},
			"sms_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignCampaignSmsMessage(),
				Optional: true,
				MaxItems: 1,
			},
			"adm_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}