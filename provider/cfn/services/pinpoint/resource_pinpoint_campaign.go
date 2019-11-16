// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-campaign.html

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointCampaignType string = "AWS::Pinpoint::Campaign"

var pinpointCampaignProperties map[string]string = map[string]string{
	"description": "Description",
	"segment_id": "SegmentId",
	"is_paused": "IsPaused",
	"additional_treatments": "AdditionalTreatments",
	"name": "Name",
	"segment_version": "SegmentVersion",
	"treatment_description": "TreatmentDescription",
	"message_configuration": "MessageConfiguration",
	"limits": "Limits",
	"holdout_percent": "HoldoutPercent",
	"schedule": "Schedule",
	"application_id": "ApplicationId",
	"campaign_hook": "CampaignHook",
	"tags": "Tags",
	"treatment_name": "TreatmentName",
}

func ResourcePinpointCampaign() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointCampaignExists,
		Read: resourcePinpointCampaignRead,
		Create: resourcePinpointCampaignCreate,
		Update: resourcePinpointCampaignUpdate,
		Delete: resourcePinpointCampaignDelete,
		CustomizeDiff: resourcePinpointCampaignCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"segment_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"is_paused": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"additional_treatments": {
				Type: schema.TypeList,
				Elem: propertyCampaignWriteTreatmentResource(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"segment_version": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"treatment_description": {
				Type: schema.TypeString,
				Optional: true,
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
				Optional: true,
				MaxItems: 1,
			},
			"holdout_percent": {
				Type: schema.TypeInt,
				Optional: true,
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
			},
			"campaign_hook": {
				Type: schema.TypeList,
				Elem: propertyCampaignCampaignHook(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"treatment_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointCampaignExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointCampaignRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointCampaignType, ResourcePinpointCampaign(), data, meta)
}

func resourcePinpointCampaignCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointCampaignType, ResourcePinpointCampaign(), data, pinpointCampaignProperties, meta)
}

func resourcePinpointCampaignUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointCampaignType, ResourcePinpointCampaign(), data, pinpointCampaignProperties, meta)
}

func resourcePinpointCampaignDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointCampaignType, data, meta)
}

func resourcePinpointCampaignCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointCampaignType, data, meta)
}
