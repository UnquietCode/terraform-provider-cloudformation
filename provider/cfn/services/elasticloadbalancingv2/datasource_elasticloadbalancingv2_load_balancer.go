// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html

package elasticloadbalancingv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticLoadBalancingV2LoadBalancerType string = "AWS::ElasticLoadBalancingV2::LoadBalancer"

func DatasourceElasticLoadBalancingV2LoadBalancer() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElasticLoadBalancingV2LoadBalancerRead,
		
		Schema: map[string]*schema.Schema{
			"ip_address_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"load_balancer_attributes": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerLoadBalancerAttribute(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
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
			"subnet_mappings": {
				Type: schema.TypeSet,
				Elem: propertyLoadBalancerSubnetMapping(),
				Optional: true,
			},
			"subnets": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"tags": misc.PropertyTags(),
			"type": {
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

func datasourceElasticLoadBalancingV2LoadBalancerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingV2LoadBalancerType, DatasourceElasticLoadBalancingV2LoadBalancer(), data, meta)
}
