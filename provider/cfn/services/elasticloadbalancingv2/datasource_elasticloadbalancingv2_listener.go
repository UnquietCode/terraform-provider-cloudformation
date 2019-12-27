// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listener.html

package elasticloadbalancingv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticLoadBalancingV2ListenerType string = "AWS::ElasticLoadBalancingV2::Listener"

func DatasourceElasticLoadBalancingV2Listener() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElasticLoadBalancingV2ListenerRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceElasticLoadBalancingV2ListenerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingV2ListenerType, DatasourceElasticLoadBalancingV2Listener(), data, meta)
}
