// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointApplicationSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointApplicationSettingsCreate,
		Read:   resourcePinpointApplicationSettingsRead,
		Update: resourcePinpointApplicationSettingsUpdate,
		Delete: resourcePinpointApplicationSettingsDelete,

		Schema: map[string]*schema.Schema{
			"quiet_time": {
				Type: schema.TypeList,
				Elem: propertyQuietTime(),
				Required: false,
				MaxItems: 1,
			},
			"limits": {
				Type: schema.TypeList,
				Elem: propertyLimits(),
				Required: false,
				MaxItems: 1,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"campaign_hook": {
				Type: schema.TypeList,
				Elem: propertyCampaignHook(),
				Required: false,
				MaxItems: 1,
			},
			"cloud_watch_metrics_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}

func resourcePinpointApplicationSettingsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::ApplicationSettings", data, meta)
}

func resourcePinpointApplicationSettingsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::ApplicationSettings", data, meta)
}

func resourcePinpointApplicationSettingsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::ApplicationSettings", data, meta)
}

func resourcePinpointApplicationSettingsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::ApplicationSettings", data, meta)
}