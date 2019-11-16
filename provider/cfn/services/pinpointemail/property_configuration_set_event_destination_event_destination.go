// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationseteventdestination-eventdestination.html

package pinpointemail

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var configurationSetEventDestinationEventDestinationProperties map[string]string = map[string]string{
	"sns_destination": "SnsDestination",
	"cloud_watch_destination": "CloudWatchDestination",
	"enabled": "Enabled",
	"matching_event_types": "MatchingEventTypes",
	"pinpoint_destination": "PinpointDestination",
	"kinesis_firehose_destination": "KinesisFirehoseDestination",
}

func propertyConfigurationSetEventDestinationEventDestination(extras...string) *schema.Resource {
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
			"sns_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationSnsDestination(),
				Optional: true,
				MaxItems: 1,
			},
			"cloud_watch_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationCloudWatchDestination(),
				Optional: true,
				MaxItems: 1,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"matching_event_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"pinpoint_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationPinpointDestination(),
				Optional: true,
				MaxItems: 1,
			},
			"kinesis_firehose_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationKinesisFirehoseDestination(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
