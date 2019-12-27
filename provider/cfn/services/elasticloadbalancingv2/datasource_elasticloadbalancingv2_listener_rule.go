// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-listenerrule.html

package elasticloadbalancingv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const elasticLoadBalancingV2ListenerRuleType string = "AWS::ElasticLoadBalancingV2::ListenerRule"

func DatasourceElasticLoadBalancingV2ListenerRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceElasticLoadBalancingV2ListenerRuleRead,
		
		Schema: map[string]*schema.Schema{
			"actions": {
				Type: schema.TypeSet,
				Elem: propertyListenerRuleAction(),
				Required: true,
			},
			"conditions": {
				Type: schema.TypeSet,
				Elem: propertyListenerRuleRuleCondition(),
				Required: true,
			},
			"listener_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"priority": {
				Type: schema.TypeInt,
				Required: true,
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

func datasourceElasticLoadBalancingV2ListenerRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(elasticLoadBalancingV2ListenerRuleType, DatasourceElasticLoadBalancingV2ListenerRule(), data, meta)
}
