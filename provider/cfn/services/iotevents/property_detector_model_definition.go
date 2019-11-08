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

func propertyDetectorModelDefinition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"states": {
				Type: schema.TypeList,
				Elem: propertyState(),
				Required: false,
			},
			"initial_state_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}