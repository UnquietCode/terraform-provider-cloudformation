// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceKinesisAnalyticsV2ApplicationReferenceDataSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceCreate,
		Read:   resourceKinesisAnalyticsV2ApplicationReferenceDataSourceRead,
		Update: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceUpdate,
		Delete: resourceKinesisAnalyticsV2ApplicationReferenceDataSourceDelete,

		Schema: map[string]*schema.Schema{
			"application_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"reference_data_source": {
				Type: schema.TypeList,
				Elem: propertyReferenceDataSource(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource", data, meta)
}

func resourceKinesisAnalyticsV2ApplicationReferenceDataSourceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::KinesisAnalyticsV2::ApplicationReferenceDataSource", data, meta)
}