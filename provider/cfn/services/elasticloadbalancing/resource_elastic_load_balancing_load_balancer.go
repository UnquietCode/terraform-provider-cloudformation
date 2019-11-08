// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package elasticloadbalancing

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticLoadBalancingLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticLoadBalancingLoadBalancerCreate,
		Read:   resourceElasticLoadBalancingLoadBalancerRead,
		Update: resourceElasticLoadBalancingLoadBalancerUpdate,
		Delete: resourceElasticLoadBalancingLoadBalancerDelete,

		Schema: map[string]*schema.Schema{
			"access_logging_policy": {
				Type: schema.TypeList,
				Elem: propertyAccessLoggingPolicy(),
				Required: false,
				MaxItems: 1,
			},
			"app_cookie_stickiness_policy": {
				Type: schema.TypeSet,
				Elem: propertyAppCookieStickinessPolicy(),
				Required: false,
			},
			"availability_zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"connection_draining_policy": {
				Type: schema.TypeList,
				Elem: propertyConnectionDrainingPolicy(),
				Required: false,
				MaxItems: 1,
			},
			"connection_settings": {
				Type: schema.TypeList,
				Elem: propertyConnectionSettings(),
				Required: false,
				MaxItems: 1,
			},
			"cross_zone": {
				Type: schema.TypeBool,
				Required: false,
			},
			"health_check": {
				Type: schema.TypeList,
				Elem: propertyHealthCheck(),
				Required: false,
				MaxItems: 1,
			},
			"instances": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"lb_cookie_stickiness_policy": {
				Type: schema.TypeSet,
				Elem: propertyLBCookieStickinessPolicy(),
				Required: false,
			},
			"listeners": {
				Type: schema.TypeSet,
				Elem: propertyListeners(),
				Required: true,
			},
			"load_balancer_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"policies": {
				Type: schema.TypeSet,
				Elem: propertyPolicies(),
				Required: false,
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
		},
	}
}

func resourceElasticLoadBalancingLoadBalancerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancing::LoadBalancer", data, meta)
}

func resourceElasticLoadBalancingLoadBalancerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancing::LoadBalancer", data, meta)
}

func resourceElasticLoadBalancingLoadBalancerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancing::LoadBalancer", data, meta)
}

func resourceElasticLoadBalancingLoadBalancerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancing::LoadBalancer", data, meta)
}