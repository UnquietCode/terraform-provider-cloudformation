// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appmesh

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyVirtualNodeHealthCheck(extras...string) *schema.Resource {
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
			"path": {
				Type: schema.TypeString,
				Required: false,
			},
			"unhealthy_threshold": {
				Type: schema.TypeInt,
				Required: true,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"healthy_threshold": {
				Type: schema.TypeInt,
				Required: true,
			},
			"timeout_millis": {
				Type: schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"interval_millis": {
				Type: schema.TypeInt,
				Required: true,
			},
		},
	}
}