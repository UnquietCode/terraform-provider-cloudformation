// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-applicationoutput-output.html

package kinesisanalyticsv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationOutputOutput(extras...string) *schema.Resource {
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
			"destination_schema": {
				Type: schema.TypeList,
				Elem: propertyApplicationOutputDestinationSchema(),
				Required: true,
				MaxItems: 1,
			},
			"lambda_output": {
				Type: schema.TypeList,
				Elem: propertyApplicationOutputLambdaOutput(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_firehose_output": {
				Type: schema.TypeList,
				Elem: propertyApplicationOutputKinesisFirehoseOutput(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_streams_output": {
				Type: schema.TypeList,
				Elem: propertyApplicationOutputKinesisStreamsOutput(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}