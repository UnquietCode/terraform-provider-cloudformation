// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-policy.html

package elasticloadbalancing

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var loadBalancerPoliciesProperties map[string]string = map[string]string{
	"attributes": "Attributes",
	"instance_ports": "InstancePorts",
	"load_balancer_ports": "LoadBalancerPorts",
	"policy_name": "PolicyName",
	"policy_type": "PolicyType",
}

func propertyLoadBalancerPolicies(extras...string) *schema.Resource {
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
			"attributes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeMap},
				Required: true,
			},
			"instance_ports": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"load_balancer_ports": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"policy_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"policy_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
