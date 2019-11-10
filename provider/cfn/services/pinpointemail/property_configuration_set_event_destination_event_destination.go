// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyConfigurationSetEventDestinationEventDestination() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"sns_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationSnsDestination(),
				Required: false,
				MaxItems: 1,
			},
			"cloud_watch_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationCloudWatchDestination(),
				Required: false,
				MaxItems: 1,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"matching_event_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"pinpoint_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationPinpointDestination(),
				Required: false,
				MaxItems: 1,
			},
			"kinesis_firehose_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationKinesisFirehoseDestination(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}