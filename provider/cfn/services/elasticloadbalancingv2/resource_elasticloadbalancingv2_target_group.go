// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html

package elasticloadbalancingv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticLoadBalancingV2TargetGroupType string = "AWS::ElasticLoadBalancingV2::TargetGroup"

func ResourceElasticLoadBalancingV2TargetGroup() *schema.Resource {
	return &schema.Resource{
		Read: resourceElasticLoadBalancingV2TargetGroupRead,
		Create: resourceElasticLoadBalancingV2TargetGroupCreate,
		Update: resourceElasticLoadBalancingV2TargetGroupUpdate,
		Delete: resourceElasticLoadBalancingV2TargetGroupDelete,
		CustomizeDiff: resourceElasticLoadBalancingV2TargetGroupCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"health_check_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"health_check_interval_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"health_check_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"health_check_port": {
				Type: schema.TypeString,
				Optional: true,
			},
			"health_check_protocol": {
				Type: schema.TypeString,
				Optional: true,
			},
			"health_check_timeout_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"healthy_threshold_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"matcher": {
				Type: schema.TypeSet,
				Elem: propertyTargetGroupMatcher(),
				Optional: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"port": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"protocol": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"target_group_attributes": {
				Type: schema.TypeSet,
				Elem: propertyTargetGroupTargetGroupAttribute(),
				Optional: true,
			},
			"target_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"targets": {
				Type: schema.TypeSet,
				Elem: propertyTargetGroupTargetDescription(),
				Optional: true,
			},
			"unhealthy_threshold_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"vpc_id": {
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

func resourceElasticLoadBalancingV2TargetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingV2TargetGroupType, ResourceElasticLoadBalancingV2TargetGroup(), data, meta)
}

func resourceElasticLoadBalancingV2TargetGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(elasticLoadBalancingV2TargetGroupType, ResourceElasticLoadBalancingV2TargetGroup(), data, meta)
}

func resourceElasticLoadBalancingV2TargetGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(elasticLoadBalancingV2TargetGroupType, ResourceElasticLoadBalancingV2TargetGroup(), data, meta)
}

func resourceElasticLoadBalancingV2TargetGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(elasticLoadBalancingV2TargetGroupType, data, meta)
}

func resourceElasticLoadBalancingV2TargetGroupCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(elasticLoadBalancingV2TargetGroupType, data, meta)
}
