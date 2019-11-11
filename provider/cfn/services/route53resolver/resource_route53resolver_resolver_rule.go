// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverrule.html

package route53resolver

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoute53ResolverResolverRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRoute53ResolverResolverRuleCreate,
		Read:   resourceRoute53ResolverResolverRuleRead,
		Update: resourceRoute53ResolverResolverRuleUpdate,
		Delete: resourceRoute53ResolverResolverRuleDelete,

		Schema: map[string]*schema.Schema{
			"resolver_endpoint_id": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resolver_rule_id": {
				Type: schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"target_ips": {
				Type: schema.TypeList,
				Elem: propertyResolverRuleTargetAddress(),
				Optional: true,
				Computed: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rule_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
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

func resourceRoute53ResolverResolverRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53Resolver::ResolverRule", ResourceRoute53ResolverResolverRule(), data, meta)
}

func resourceRoute53ResolverResolverRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53Resolver::ResolverRule", ResourceRoute53ResolverResolverRule(), data, meta)
}

func resourceRoute53ResolverResolverRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53Resolver::ResolverRule", ResourceRoute53ResolverResolverRule(), data, meta)
}

func resourceRoute53ResolverResolverRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53Resolver::ResolverRule", data, meta)
}