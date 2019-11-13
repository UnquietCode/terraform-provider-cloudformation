// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationrecorder.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigConfigurationRecorder() *schema.Resource {
	return &schema.Resource{
		Exists: resourceConfigConfigurationRecorderExists,
		Read:   resourceConfigConfigurationRecorderRead,
		Create: resourceConfigConfigurationRecorderCreate,
		Update: resourceConfigConfigurationRecorderUpdate,
		Delete: resourceConfigConfigurationRecorderDelete,
		CustomizeDiff: resourceConfigConfigurationRecorderCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"recording_group": {
				Type: schema.TypeList,
				Elem: propertyConfigurationRecorderRecordingGroup(),
				Optional: true,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConfigConfigurationRecorderExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceConfigConfigurationRecorderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigurationRecorder", ResourceConfigConfigurationRecorder(), data, meta)
}

func resourceConfigConfigurationRecorderCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigurationRecorder", ResourceConfigConfigurationRecorder(), data, meta)
}

func resourceConfigConfigurationRecorderUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigurationRecorder", ResourceConfigConfigurationRecorder(), data, meta)
}

func resourceConfigConfigurationRecorderDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigurationRecorder", data, meta)
}

func resourceConfigConfigurationRecorderCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}