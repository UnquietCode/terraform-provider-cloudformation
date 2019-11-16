// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationreferencedatasource-referencedatasource.html

package kinesisanalyticsv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var applicationReferenceDataSourceReferenceDataSourceProperties map[string]string = map[string]string{
	"reference_schema": "ReferenceSchema",
	"table_name": "TableName",
	"s3_reference_data_source": "S3ReferenceDataSource",
}

func propertyApplicationReferenceDataSourceReferenceDataSource(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reference_schema": {
				Type: schema.TypeSet,
				Elem: propertyApplicationReferenceDataSourceReferenceSchema(),
				Required: true,
				MaxItems: 1,
			},
			"table_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"s3_reference_data_source": {
				Type: schema.TypeSet,
				Elem: propertyApplicationReferenceDataSourceS3ReferenceDataSource(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
