// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationInput() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name_prefix": {
				Type: schema.TypeString,
				Required: true,
			},
			"input_schema": {
				Type: schema.TypeList,
				Elem: propertyApplicationInputSchema(),
				Required: true,
				MaxItems: 1,
			},
			"kinesis_streams_input": {
				Type: schema.TypeList,
				Elem: propertyApplicationKinesisStreamsInput(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_firehose_input": {
				Type: schema.TypeList,
				Elem: propertyApplicationKinesisFirehoseInput(),
				Required: false,
				MaxItems: 1,
			},
			"input_processing_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationInputProcessingConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"input_parallelism": {
				Type: schema.TypeList,
				Elem: propertyApplicationInputParallelism(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}