// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-targetgroup.html

package elasticloadbalancingv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticLoadBalancingV2TargetGroupType string = "AWS::ElasticLoadBalancingV2::TargetGroup"

func DatasourceElasticLoadBalancingV2TargetGroup() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElasticLoadBalancingV2TargetGroupRead,
		
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
				Type: schema.TypeList,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceElasticLoadBalancingV2TargetGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingV2TargetGroupType, DatasourceElasticLoadBalancingV2TargetGroup(), data, meta)
}
