// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-eventdestination.html

package ses

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

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
			"cloud_watch_destination": {
				Type: schema.TypeSet,
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
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kinesis_firehose_destination": {
				Type: schema.TypeSet,
				Elem: propertyConfigurationSetEventDestinationKinesisFirehoseDestination(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
