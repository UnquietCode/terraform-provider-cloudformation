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

func propertyOutput() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"destination_schema": {
				Type: schema.TypeList,
				Elem: propertyDestinationSchema(),
				Required: true,
				MaxItems: 1,
			},
			"lambda_output": {
				Type: schema.TypeList,
				Elem: propertyLambdaOutput(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_firehose_output": {
				Type: schema.TypeList,
				Elem: propertyKinesisFirehoseOutput(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_streams_output": {
				Type: schema.TypeList,
				Elem: propertyKinesisStreamsOutput(),
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