// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCampaign() *schema.Resource {
	return &schema.Resource{
		Create: resourceCampaignCreate,
		Read:   resourceCampaignRead,
		Update: resourceCampaignUpdate,
		Delete: resourceCampaignDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"segment_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"is_paused": {
				Type: schema.TypeBool,
				Required: false,
			},
			"additional_treatments": {
				Type: schema.TypeList,
				Elem: propertyCampaignWriteTreatmentResource(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"segment_version": {
				Type: schema.TypeInt,
				Required: false,
			},
			"treatment_description": {
				Type: schema.TypeString,
				Required: false,
			},
			"message_configuration": {
				Type: schema.TypeList,
				Elem: propertyCampaignMessageConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"limits": {
				Type: schema.TypeList,
				Elem: propertyCampaignLimits(),
				Required: false,
				MaxItems: 1,
			},
			"holdout_percent": {
				Type: schema.TypeInt,
				Required: false,
			},
			"schedule": {
				Type: schema.TypeList,
				Elem: propertyCampaignSchedule(),
				Required: true,
				MaxItems: 1,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"campaign_hook": {
				Type: schema.TypeList,
				Elem: propertyCampaignCampaignHook(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"treatment_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceCampaignCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::Campaign", data, meta)
}

func resourceCampaignRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::Campaign", data, meta)
}

func resourceCampaignUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::Campaign", data, meta)
}

func resourceCampaignDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::Campaign", data, meta)
}