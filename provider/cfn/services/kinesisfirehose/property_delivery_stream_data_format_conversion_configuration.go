// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamDataFormatConversionConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"input_format_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamInputFormatConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"output_format_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamOutputFormatConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"schema_configuration": {
				Type: schema.TypeList,
				Elem: propertyDeliveryStreamSchemaConfiguration(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}