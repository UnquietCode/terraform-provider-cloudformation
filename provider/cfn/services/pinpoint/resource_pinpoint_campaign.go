// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointCampaign() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointCampaignCreate,
		Read:   resourcePinpointCampaignRead,
		Update: resourcePinpointCampaignUpdate,
		Delete: resourcePinpointCampaignDelete,

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

func resourcePinpointCampaignCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::Campaign", data, meta)
}

func resourcePinpointCampaignRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::Campaign", data, meta)
}

func resourcePinpointCampaignUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::Campaign", data, meta)
}

func resourcePinpointCampaignDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::Campaign", data, meta)
}