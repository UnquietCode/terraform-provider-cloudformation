// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticLoadBalancingV2ListenerType string = "AWS::ElasticLoadBalancingV2::Listener"

var elasticLoadBalancingV2ListenerProperties map[string]string = map[string]string{
	"certificates": "Certificates",
	"default_actions": "DefaultActions",
	"load_balancer_arn": "LoadBalancerArn",
	"port": "Port",
	"protocol": "Protocol",
	"ssl_policy": "SslPolicy",
}

func ResourceElasticLoadBalancingV2Listener() *schema.Resource {
	return &schema.Resource{
		Exists: resourceElasticLoadBalancingV2ListenerExists,
		Read: resourceElasticLoadBalancingV2ListenerRead,
		Create: resourceElasticLoadBalancingV2ListenerCreate,
		Update: resourceElasticLoadBalancingV2ListenerUpdate,
		Delete: resourceElasticLoadBalancingV2ListenerDelete,
		CustomizeDiff: resourceElasticLoadBalancingV2ListenerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"certificates": {
				Type: schema.TypeSet,
				Elem: propertyListenerCertificate(),
				Optional: true,
			},
			"default_actions": {
				Type: schema.TypeSet,
				Elem: propertyListenerAction(),
				Required: true,
			},
			"load_balancer_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"port": {
				Type: schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"ssl_policy": {
				Type: schema.TypeString,
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

func resourceElasticLoadBalancingV2ListenerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceElasticLoadBalancingV2ListenerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingV2ListenerType, ResourceElasticLoadBalancingV2Listener(), data, meta)
}

func resourceElasticLoadBalancingV2ListenerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elasticLoadBalancingV2ListenerType, ResourceElasticLoadBalancingV2Listener(), data, elasticLoadBalancingV2ListenerProperties, meta)
}

func resourceElasticLoadBalancingV2ListenerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elasticLoadBalancingV2ListenerType, ResourceElasticLoadBalancingV2Listener(), data, elasticLoadBalancingV2ListenerProperties, meta)
}

func resourceElasticLoadBalancingV2ListenerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elasticLoadBalancingV2ListenerType, data, meta)
}

func resourceElasticLoadBalancingV2ListenerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elasticLoadBalancingV2ListenerType, data, meta)
}
