// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-dataformatconversionconfiguration.html

package kinesisfirehose

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeliveryStreamDataFormatConversionConfiguration(extras...string) *schema.Resource {
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