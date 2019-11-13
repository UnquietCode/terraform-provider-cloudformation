// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationreferencedatasource.html

package kinesisanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisAnalyticsApplicationReferenceDataSource() *schema.Resource {
	return &schema.Resource{
		Exists: resourceKinesisAnalyticsApplicationReferenceDataSourceExists,
		Read:   resourceKinesisAnalyticsApplicationReferenceDataSourceRead,
		Create: resourceKinesisAnalyticsApplicationReferenceDataSourceCreate,
		Update: resourceKinesisAnalyticsApplicationReferenceDataSourceUpdate,
		Delete: resourceKinesisAnalyticsApplicationReferenceDataSourceDelete,
		CustomizeDiff: resourceKinesisAnalyticsApplicationReferenceDataSourceCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"reference_data_source": {
				Type: schema.TypeList,
				Elem: propertyApplicationReferenceDataSourceReferenceDataSource(),
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

func resourceKinesisAnalyticsApplicationReferenceDataSourceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKinesisAnalyticsApplicationReferenceDataSourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalytics::ApplicationReferenceDataSource", ResourceKinesisAnalyticsApplicationReferenceDataSource(), data, meta)
}

func resourceKinesisAnalyticsApplicationReferenceDataSourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalytics::ApplicationReferenceDataSource", ResourceKinesisAnalyticsApplicationReferenceDataSource(), data, meta)
}

func resourceKinesisAnalyticsApplicationReferenceDataSourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalytics::ApplicationReferenceDataSource", ResourceKinesisAnalyticsApplicationReferenceDataSource(), data, meta)
}

func resourceKinesisAnalyticsApplicationReferenceDataSourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalytics::ApplicationReferenceDataSource", data, meta)
}

func resourceKinesisAnalyticsApplicationReferenceDataSourceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}