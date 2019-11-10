// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalytics

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationReferenceDataSourceReferenceDataSource(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reference_schema": {
				Type: schema.TypeList,
				Elem: propertyApplicationReferenceDataSourceReferenceSchema(),
				Required: true,
				MaxItems: 1,
			},
			"table_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"s3_reference_data_source": {
				Type: schema.TypeList,
				Elem: propertyApplicationReferenceDataSourceS3ReferenceDataSource(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}