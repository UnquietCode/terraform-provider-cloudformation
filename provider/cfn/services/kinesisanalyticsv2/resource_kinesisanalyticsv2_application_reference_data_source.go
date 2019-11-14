// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationreferencedatasource.html

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisAnalyticsV2ApplicationReferenceDataSource() *schema.Resource {
	return &schema.Resource{
		Exists: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceExists,
		Read: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceRead,
		Create: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceCreate,
		Update: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceUpdate,
		Delete: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceDelete,
		CustomizeDiff: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceCustomizeDiff,
		
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

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource", ResourceKinesisAnalyticsV2ApplicationReferenceDataSource(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource", ResourceKinesisAnalyticsV2ApplicationReferenceDataSource(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource", ResourceKinesisAnalyticsV2ApplicationReferenceDataSource(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
