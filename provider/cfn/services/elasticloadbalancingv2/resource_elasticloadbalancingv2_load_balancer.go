// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticLoadBalancingV2LoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticLoadBalancingV2LoadBalancerCreate,
		Read:   resourceElasticLoadBalancingV2LoadBalancerRead,
		Update: resourceElasticLoadBalancingV2LoadBalancerUpdate,
		Delete: resourceElasticLoadBalancingV2LoadBalancerDelete,

		Schema: map[string]*schema.Schema{
			"canonical_hosted_zone_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"dns_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"load_balancer_full_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"load_balancer_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
				Set: schema.HashString,
			},
			"ip_address_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"load_balancer_attributes": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerLoadBalancerAttribute(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scheme": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"subnet_mappings": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerSubnetMapping(),
				Optional: true,
				ForceNew: true,
			},
			"subnets": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElasticLoadBalancingV2LoadBalancerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancingV2::LoadBalancer", data, meta)
}

func resourceElasticLoadBalancingV2LoadBalancerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancingV2::LoadBalancer", data, meta)
}

func resourceElasticLoadBalancingV2LoadBalancerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancingV2::LoadBalancer", data, meta)
}

func resourceElasticLoadBalancingV2LoadBalancerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancingV2::LoadBalancer", data, meta)
}