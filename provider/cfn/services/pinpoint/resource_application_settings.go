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

func ResourceApplicationSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationSettingsCreate,
		Read:   resourceApplicationSettingsRead,
		Update: resourceApplicationSettingsUpdate,
		Delete: resourceApplicationSettingsDelete,

		Schema: map[string]*schema.Schema{
			"quiet_time": {
				Type: schema.TypeList,
				Elem: propertyApplicationSettingsQuietTime(),
				Required: false,
				MaxItems: 1,
			},
			"limits": {
				Type: schema.TypeList,
				Elem: propertyApplicationSettingsLimits(),
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
				Elem: propertyApplicationSettingsCampaignHook(),
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

func resourceApplicationSettingsCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Pinpoint::ApplicationSettings", data, meta)
}

func resourceApplicationSettingsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Pinpoint::ApplicationSettings", data, meta)
}

func resourceApplicationSettingsUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Pinpoint::ApplicationSettings", data, meta)
}

func resourceApplicationSettingsDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Pinpoint::ApplicationSettings", data, meta)
}