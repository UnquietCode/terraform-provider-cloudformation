// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInput() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name_prefix": {
				Type: schema.TypeString,
				Required: true,
			},
			"input_schema": {
				Type: schema.TypeList,
				Elem: propertyInputSchema(),
				Required: true,
				MaxItems: 1,
			},
			"kinesis_streams_input": {
				Type: schema.TypeList,
				Elem: propertyKinesisStreamsInput(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_firehose_input": {
				Type: schema.TypeList,
				Elem: propertyKinesisFirehoseInput(),
				Required: false,
				MaxItems: 1,
			},
			"input_processing_configuration": {
				Type: schema.TypeList,
				Elem: propertyInputProcessingConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"input_parallelism": {
				Type: schema.TypeList,
				Elem: propertyInputParallelism(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}