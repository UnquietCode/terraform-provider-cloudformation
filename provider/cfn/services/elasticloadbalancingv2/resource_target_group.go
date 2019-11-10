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

func ResourceTargetGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTargetGroupCreate,
		Read:   resourceTargetGroupRead,
		Update: resourceTargetGroupUpdate,
		Delete: resourceTargetGroupDelete,

		Schema: map[string]*schema.Schema{
			"health_check_enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"health_check_interval_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"health_check_path": {
				Type: schema.TypeString,
				Required: false,
			},
			"health_check_port": {
				Type: schema.TypeString,
				Required: false,
			},
			"health_check_protocol": {
				Type: schema.TypeString,
				Required: false,
			},
			"health_check_timeout_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"healthy_threshold_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"matcher": {
				Type: schema.TypeList,
				Elem: propertyTargetGroupMatcher(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"protocol": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"target_group_attributes": {
				Type: schema.TypeSet,
				Elem: propertyTargetGroupTargetGroupAttribute(),
				Required: false,
			},
			"target_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"targets": {
				Type: schema.TypeSet,
				Elem: propertyTargetGroupTargetDescription(),
				Required: false,
			},
			"unhealthy_threshold_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceTargetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ElasticLoadBalancingV2::TargetGroup", data, meta)
}

func resourceTargetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ElasticLoadBalancingV2::TargetGroup", data, meta)
}

func resourceTargetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ElasticLoadBalancingV2::TargetGroup", data, meta)
}

func resourceTargetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ElasticLoadBalancingV2::TargetGroup", data, meta)
}