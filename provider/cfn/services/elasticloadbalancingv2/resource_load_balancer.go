// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceLoadBalancerCreate,
		Read:   resourceLoadBalancerRead,
		Update: resourceLoadBalancerUpdate,
		Delete: resourceLoadBalancerDelete,

		Schema: map[string]*schema.Schema{
			"ip_address_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"load_balancer_attributes": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerLoadBalancerAttribute(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"scheme": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"subnet_mappings": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerSubnetMapping(),
				Required: false,
				ForceNew: true,
			},
			"subnets": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceLoadBalancerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancingV2::LoadBalancer", data, meta)
}

func resourceLoadBalancerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancingV2::LoadBalancer", data, meta)
}

func resourceLoadBalancerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancingV2::LoadBalancer", data, meta)
}

func resourceLoadBalancerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancingV2::LoadBalancer", data, meta)
}