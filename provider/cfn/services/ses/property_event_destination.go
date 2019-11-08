// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEventDestination() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_watch_destination": {
				Type: schema.TypeList,
				Elem: propertyCloudWatchDestination(),
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
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
			"kinesis_firehose_destination": {
				Type: schema.TypeList,
				Elem: propertyKinesisFirehoseDestination(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}