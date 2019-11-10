// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticloadbalancing

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLoadBalancerListeners(extras...string) *schema.Resource {
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
			"instance_port": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_protocol": {
				Type: schema.TypeString,
				Required: false,
			},
			"load_balancer_port": {
				Type: schema.TypeString,
				Required: true,
			},
			"policy_names": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"ssl_certificate_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}