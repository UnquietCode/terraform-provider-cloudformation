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

func propertyReferenceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"record_encoding": {
				Type: schema.TypeString,
				Required: false,
			},
			"record_columns": {
				Type: schema.TypeList,
				Elem: propertyRecordColumn(),
				Required: true,
			},
			"record_format": {
				Type: schema.TypeList,
				Elem: propertyRecordFormat(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}