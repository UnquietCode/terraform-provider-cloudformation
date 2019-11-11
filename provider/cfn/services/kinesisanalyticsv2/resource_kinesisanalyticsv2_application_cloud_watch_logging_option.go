// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionCreate,
		Read:   resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionRead,
		Update: resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionUpdate,
		Delete: resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionDelete,

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
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption", ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption", ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption", ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption", data, meta)
}