// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationInputSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"record_encoding": {
				Type: schema.TypeString,
				Required: false,
			},
			"record_columns": {
				Type: schema.TypeList,
				Elem: propertyApplicationRecordColumn(),
				Required: true,
			},
			"record_format": {
				Type: schema.TypeList,
				Elem: propertyApplicationRecordFormat(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}