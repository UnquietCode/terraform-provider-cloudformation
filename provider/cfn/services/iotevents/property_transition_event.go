// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iotevents

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTransitionEvent() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"condition": {
				Type: schema.TypeString,
				Required: false,
			},
			"actions": {
				Type: schema.TypeList,
				Elem: propertyAction(),
				Required: false,
			},
			"next_state": {
				Type: schema.TypeString,
				Required: false,
			},
			"event_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}