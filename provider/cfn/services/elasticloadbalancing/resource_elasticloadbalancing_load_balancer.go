// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html

package elasticloadbalancing

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticLoadBalancingLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElasticLoadBalancingLoadBalancerExists,
		Read: resourceElasticLoadBalancingLoadBalancerRead,
		Create: resourceElasticLoadBalancingLoadBalancerCreate,
		Update: resourceElasticLoadBalancingLoadBalancerUpdate,
		Delete: resourceElasticLoadBalancingLoadBalancerDelete,
		CustomizeDiff: resourceElasticLoadBalancingLoadBalancerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"access_logging_policy": {
				Type: schema.TypeList,
				Elem: propertyLoadBalancerAccessLoggingPolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"app_cookie_stickiness_policy": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerAppCookieStickinessPolicy(),
				Optional: true,
			},
			"availability_zones": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"connection_draining_policy": {
				Type: schema.TypeList,
				Elem: propertyLoadBalancerConnectionDrainingPolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"connection_settings": {
				Type: schema.TypeList,
				Elem: propertyLoadBalancerConnectionSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"cross_zone": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"health_check": {
				Type: schema.TypeList,
				Elem: propertyLoadBalancerHealthCheck(),
				Optional: true,
				MaxItems: 1,
			},
			"instances": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"lb_cookie_stickiness_policy": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerLBCookieStickinessPolicy(),
				Optional: true,
			},
			"listeners": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerListeners(),
				Required: true,
			},
			"load_balancer_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"policies": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerPolicies(),
				Optional: true,
			},
			"scheme": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElasticLoadBalancingLoadBalancerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElasticLoadBalancingLoadBalancerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancing::LoadBalancer", ResourceElasticLoadBalancingLoadBalancer(), data, meta)
}

func resourceElasticLoadBalancingLoadBalancerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancing::LoadBalancer", ResourceElasticLoadBalancingLoadBalancer(), data, meta)
}

func resourceElasticLoadBalancingLoadBalancerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancing::LoadBalancer", ResourceElasticLoadBalancingLoadBalancer(), data, meta)
}

func resourceElasticLoadBalancingLoadBalancerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancing::LoadBalancer", data, meta)
}

func resourceElasticLoadBalancingLoadBalancerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

