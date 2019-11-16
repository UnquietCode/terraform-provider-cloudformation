// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-input.html

package kinesisanalyticsv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationInput(extras...string) *schema.Resource {
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
			"name_prefix": {
				Type: schema.TypeString,
				Required: true,
			},
			"input_schema": {
				Type: schema.TypeSet,
				Elem: propertyApplicationInputSchema(),
				Required: true,
				MaxItems: 1,
			},
			"kinesis_streams_input": {
				Type: schema.TypeSet,
				Elem: propertyApplicationKinesisStreamsInput(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis_firehose_input": {
				Type: schema.TypeSet,
				Elem: propertyApplicationKinesisFirehoseInput(),
				Optional: true,
				MaxItems: 1,
			},
			"input_processing_configuration": {
				Type: schema.TypeSet,
				Elem: propertyApplicationInputProcessingConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"input_parallelism": {
				Type: schema.TypeSet,
				Elem: propertyApplicationInputParallelism(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
