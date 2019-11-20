// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
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

const elasticLoadBalancingLoadBalancerType string = "AWS::ElasticLoadBalancing::LoadBalancer"

func ResourceElasticLoadBalancingLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Read: resourceElasticLoadBalancingLoadBalancerRead,
		Create: resourceElasticLoadBalancingLoadBalancerCreate,
		Update: resourceElasticLoadBalancingLoadBalancerUpdate,
		Delete: resourceElasticLoadBalancingLoadBalancerDelete,
		CustomizeDiff: resourceElasticLoadBalancingLoadBalancerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"access_logging_policy": {
				Type: schema.TypeSet,
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
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerConnectionDrainingPolicy(),
				Optional: true,
				MaxItems: 1,
			},
			"connection_settings": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerConnectionSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"cross_zone": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"health_check": {
				Type: schema.TypeSet,
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
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceElasticLoadBalancingLoadBalancerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingLoadBalancerType, ResourceElasticLoadBalancingLoadBalancer(), data, meta)
}

func resourceElasticLoadBalancingLoadBalancerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elasticLoadBalancingLoadBalancerType, ResourceElasticLoadBalancingLoadBalancer(), data, meta)
}

func resourceElasticLoadBalancingLoadBalancerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elasticLoadBalancingLoadBalancerType, ResourceElasticLoadBalancingLoadBalancer(), data, meta)
}

func resourceElasticLoadBalancingLoadBalancerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elasticLoadBalancingLoadBalancerType, data, meta)
}

func resourceElasticLoadBalancingLoadBalancerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elasticLoadBalancingLoadBalancerType, data, meta)
}
