// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceListener() *schema.Resource {
	return &schema.Resource{
		Create: resourceListenerCreate,
		Read:   resourceListenerRead,
		Update: resourceListenerUpdate,
		Delete: resourceListenerDelete,

		Schema: map[string]*schema.Schema{
			"certificates": {
				Type: schema.TypeSet,
				Elem: propertyListenerCertificate(),
				Required: false,
			},
			"default_actions": {
				Type: schema.TypeSet,
				Elem: propertyListenerAction(),
				Required: true,
			},
			"load_balancer_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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
				Required: false,
			},
		},
	}
}

func resourceListenerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancingV2::Listener", data, meta)
}

func resourceListenerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancingV2::Listener", data, meta)
}

func resourceListenerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancingV2::Listener", data, meta)
}

func resourceListenerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancingV2::Listener", data, meta)
}