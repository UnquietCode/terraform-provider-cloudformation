// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverrule.html

package route53resolver

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53ResolverResolverRuleType string = "AWS::Route53Resolver::ResolverRule"

func ResourceRoute53ResolverResolverRule() *schema.Resource {
	return &schema.Resource{
		Read: resourceRoute53ResolverResolverRuleRead,
		Create: resourceRoute53ResolverResolverRuleCreate,
		Update: resourceRoute53ResolverResolverRuleUpdate,
		Delete: resourceRoute53ResolverResolverRuleDelete,
		CustomizeDiff: resourceRoute53ResolverResolverRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"resolver_endpoint_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"rule_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_ips": {
				Type: schema.TypeList,
				Elem: propertyResolverRuleTargetAddress(),
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceRoute53ResolverResolverRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53ResolverResolverRuleType, ResourceRoute53ResolverResolverRule(), data, meta)
}

func resourceRoute53ResolverResolverRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(route53ResolverResolverRuleType, ResourceRoute53ResolverResolverRule(), data, meta)
}

func resourceRoute53ResolverResolverRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(route53ResolverResolverRuleType, ResourceRoute53ResolverResolverRule(), data, meta)
}

func resourceRoute53ResolverResolverRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(route53ResolverResolverRuleType, data, meta)
}

func resourceRoute53ResolverResolverRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(route53ResolverResolverRuleType, data, meta)
}
