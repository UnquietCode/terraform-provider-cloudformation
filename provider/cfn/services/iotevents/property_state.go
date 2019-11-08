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

func propertyState() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"on_input": {
				Type: schema.TypeList,
				Elem: propertyOnInput(),
				Required: false,
				MaxItems: 1,
			},
			"on_exit": {
				Type: schema.TypeList,
				Elem: propertyOnExit(),
				Required: false,
				MaxItems: 1,
			},
			"state_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"on_enter": {
				Type: schema.TypeList,
				Elem: propertyOnEnter(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}