// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotanalytics

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDatasetContentDeliveryRuleDestination() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"iot_events_destination_configuration": {
				Type: schema.TypeList,
				Elem: propertyIotEventsDestinationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"s3_destination_configuration": {
				Type: schema.TypeList,
				Elem: propertyS3DestinationConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}