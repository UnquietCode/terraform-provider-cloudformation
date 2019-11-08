// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesisanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyReferenceDataSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reference_schema": {
				Type: schema.TypeList,
				Elem: propertyReferenceSchema(),
				Required: true,
				MaxItems: 1,
			},
			"table_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"s3_reference_data_source": {
				Type: schema.TypeList,
				Elem: propertyS3ReferenceDataSource(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}