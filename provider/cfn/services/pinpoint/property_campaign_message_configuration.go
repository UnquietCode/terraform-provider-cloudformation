// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCampaignMessageConfiguration() *schema.Resource {
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