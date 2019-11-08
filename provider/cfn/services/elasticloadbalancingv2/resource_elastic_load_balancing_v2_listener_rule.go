// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceElasticLoadBalancingV2ListenerRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceElasticLoadBalancingV2ListenerRuleCreate,
		Read:   resourceElasticLoadBalancingV2ListenerRuleRead,
		Update: resourceElasticLoadBalancingV2ListenerRuleUpdate,
		Delete: resourceElasticLoadBalancingV2ListenerRuleDelete,

		Schema: map[string]*schema.Schema{
			"actions": {
				Type: schema.TypeSet,
				Elem: propertyAction(),
				Required: true,
			},
			"conditions": {
				Type: schema.TypeSet,
				Elem: propertyRuleCondition(),
				Required: true,
			},
			"listener_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"priority": {
				Type: schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceElasticLoadBalancingV2ListenerRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancingV2::ListenerRule", data, meta)
}

func resourceElasticLoadBalancingV2ListenerRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancingV2::ListenerRule", data, meta)
}

func resourceElasticLoadBalancingV2ListenerRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancingV2::ListenerRule", data, meta)
}

func resourceElasticLoadBalancingV2ListenerRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancingV2::ListenerRule", data, meta)
}