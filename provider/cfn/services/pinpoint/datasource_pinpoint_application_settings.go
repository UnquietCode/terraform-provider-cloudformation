// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-applicationsettings.html

package pinpoint

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointApplicationSettingsType string = "AWS::Pinpoint::ApplicationSettings"

func DatasourcePinpointApplicationSettings() *schema.Resource {
	return &schema.Resource{
		Read: datasourcePinpointApplicationSettingsRead,
		
		Schema: map[string]*schema.Schema{
			"quiet_time": {
				Type: schema.TypeList,
				Elem: propertyApplicationSettingsQuietTime(),
				Optional: true,
				MaxItems: 1,
			},
			"limits": {
				Type: schema.TypeList,
				Elem: propertyApplicationSettingsLimits(),
				Optional: true,
				MaxItems: 1,
			},
			"application_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"campaign_hook": {
				Type: schema.TypeList,
				Elem: propertyApplicationSettingsCampaignHook(),
				Optional: true,
				MaxItems: 1,
			},
			"cloud_watch_metrics_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourcePinpointApplicationSettingsRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointApplicationSettingsType, DatasourcePinpointApplicationSettings(), data, meta)
}
