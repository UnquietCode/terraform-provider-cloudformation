// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceConfigConfigurationRecorderCreate,
		Read:   resourceConfigConfigurationRecorderRead,
		Update: resourceConfigConfigurationRecorderUpdate,
		Delete: resourceConfigConfigurationRecorderDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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

func resourceConfigConfigurationRecorderCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigurationRecorder", ResourceConfigConfigurationRecorder(), data, meta)
}

func resourceConfigConfigurationRecorderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigurationRecorder", ResourceConfigConfigurationRecorder(), data, meta)
}

func resourceConfigConfigurationRecorderUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigurationRecorder", ResourceConfigConfigurationRecorder(), data, meta)
}

func resourceConfigConfigurationRecorderDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigurationRecorder", data, meta)
}