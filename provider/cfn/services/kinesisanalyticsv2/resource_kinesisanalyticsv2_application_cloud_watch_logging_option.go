// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsV2ApplicationCloudWatchLoggingOptionType string = "AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption"

func ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption() *schema.Resource {
	return &schema.Resource{
		Read: resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionRead,
		Create: resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionCreate,
		Update: resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionUpdate,
		Delete: resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionDelete,
		CustomizeDiff: resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"cloud_watch_logging_option": {
				Type: schema.TypeSet,
				Elem: propertyApplicationCloudWatchLoggingOptionCloudWatchLoggingOption(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsV2ApplicationCloudWatchLoggingOptionType, ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kinesisAnalyticsV2ApplicationCloudWatchLoggingOptionType, ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kinesisAnalyticsV2ApplicationCloudWatchLoggingOptionType, ResourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kinesisAnalyticsV2ApplicationCloudWatchLoggingOptionType, data, meta)
}

func resourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kinesisAnalyticsV2ApplicationCloudWatchLoggingOptionType, data, meta)
}
