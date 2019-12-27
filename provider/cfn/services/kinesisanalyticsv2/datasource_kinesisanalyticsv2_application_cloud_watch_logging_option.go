// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationcloudwatchloggingoption.html

package kinesisanalyticsv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsV2ApplicationCloudWatchLoggingOptionType string = "AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption"

func DatasourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption() *schema.Resource {
	return &schema.Resource{
		Read: datasourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionRead,
		
		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"cloud_watch_logging_option": {
				Type: schema.TypeList,
				Elem: propertyApplicationCloudWatchLoggingOptionCloudWatchLoggingOption(),
				Required: true,
				MaxItems: 1,
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

func datasourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(kinesisAnalyticsV2ApplicationCloudWatchLoggingOptionType, DatasourceKinesisAnalyticsV2ApplicationCloudWatchLoggingOption(), data, meta)
}
