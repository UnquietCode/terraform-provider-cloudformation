// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"apns_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Required: false,
				MaxItems: 1,
			},
			"baidu_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Required: false,
				MaxItems: 1,
			},
			"default_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Required: false,
				MaxItems: 1,
			},
			"email_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignCampaignEmailMessage(),
				Required: false,
				MaxItems: 1,
			},
			"gcm_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Required: false,
				MaxItems: 1,
			},
			"sms_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignCampaignSmsMessage(),
				Required: false,
				MaxItems: 1,
			},
			"adm_message": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessage(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}