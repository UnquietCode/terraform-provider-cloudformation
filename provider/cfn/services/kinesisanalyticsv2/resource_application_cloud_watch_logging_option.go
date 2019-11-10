// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApplicationCloudWatchLoggingOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCloudWatchLoggingOptionCreate,
		Read:   resourceApplicationCloudWatchLoggingOptionRead,
		Update: resourceApplicationCloudWatchLoggingOptionUpdate,
		Delete: resourceApplicationCloudWatchLoggingOptionDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cloud_watch_logging_option": {
				Type: schema.TypeList,
				Elem: propertyApplicationCloudWatchLoggingOptionCloudWatchLoggingOption(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceApplicationCloudWatchLoggingOptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption", data, meta)
}

func resourceApplicationCloudWatchLoggingOptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption", data, meta)
}

func resourceApplicationCloudWatchLoggingOptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption", data, meta)
}

func resourceApplicationCloudWatchLoggingOptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption", data, meta)
}