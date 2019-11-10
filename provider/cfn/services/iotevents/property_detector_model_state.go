// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iotevents

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDetectorModelState(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"on_input": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelOnInput(),
				Required: false,
				MaxItems: 1,
			},
			"on_exit": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelOnExit(),
				Required: false,
				MaxItems: 1,
			},
			"state_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"on_enter": {
				Type: schema.TypeList,
				Elem: propertyDetectorModelOnEnter(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}