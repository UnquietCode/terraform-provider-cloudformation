// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigurationRecorder() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationRecorderCreate,
		Read:   resourceConfigurationRecorderRead,
		Update: resourceConfigurationRecorderUpdate,
		Delete: resourceConfigurationRecorderDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"recording_group": {
				Type: schema.TypeList,
				Elem: propertyConfigurationRecorderRecordingGroup(),
				Required: false,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConfigurationRecorderCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigurationRecorder", data, meta)
}

func resourceConfigurationRecorderRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigurationRecorder", data, meta)
}

func resourceConfigurationRecorderUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigurationRecorder", data, meta)
}

func resourceConfigurationRecorderDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigurationRecorder", data, meta)
}