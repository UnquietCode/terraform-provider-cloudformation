// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalyticsv2-applicationreferencedatasource.html

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const kinesisAnalyticsV2ApplicationReferenceDataSourceType string = "AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource"

var kinesisAnalyticsV2ApplicationReferenceDataSourceProperties map[string]string = map[string]string{
	"application_name": "ApplicationName",
	"reference_data_source": "ReferenceDataSource",
}

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
	return plugin.ResourceRead(kinesisAnalyticsV2ApplicationReferenceDataSourceType, ResourceKinesisAnalyticsV2ApplicationReferenceDataSource(), data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(kinesisAnalyticsV2ApplicationReferenceDataSourceType, ResourceKinesisAnalyticsV2ApplicationReferenceDataSource(), data, kinesisAnalyticsV2ApplicationReferenceDataSourceProperties, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(kinesisAnalyticsV2ApplicationReferenceDataSourceType, ResourceKinesisAnalyticsV2ApplicationReferenceDataSource(), data, kinesisAnalyticsV2ApplicationReferenceDataSourceProperties, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(kinesisAnalyticsV2ApplicationReferenceDataSourceType, data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(kinesisAnalyticsV2ApplicationReferenceDataSourceType, data, meta)
}
