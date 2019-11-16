// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-firehose.html

package iotevents

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var detectorModelFirehoseProperties map[string]string = map[string]string{
	"delivery_stream_name": "DeliveryStreamName",
	"separator": "Separator",
}

func propertyDetectorModelFirehose(extras...string) *schema.Resource {
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
			"delivery_stream_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"separator": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
