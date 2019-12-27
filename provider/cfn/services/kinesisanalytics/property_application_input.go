// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html

package kinesisanalytics

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
				Type: schema.TypeList,
				Elem: propertyApplicationInputSchema(),
				Required: true,
				MaxItems: 1,
			},
			"kinesis_streams_input": {
				Type: schema.TypeList,
				Elem: propertyApplicationKinesisStreamsInput(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis_firehose_input": {
				Type: schema.TypeList,
				Elem: propertyApplicationKinesisFirehoseInput(),
				Optional: true,
				MaxItems: 1,
			},
			"input_processing_configuration": {
				Type: schema.TypeList,
				Elem: propertyApplicationInputProcessingConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"input_parallelism": {
				Type: schema.TypeList,
				Elem: propertyApplicationInputParallelism(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
