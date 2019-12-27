// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceRoute53ResolverResolverRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceRoute53ResolverResolverRuleRead,
		
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceRoute53ResolverResolverRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53ResolverResolverRuleType, DatasourceRoute53ResolverResolverRule(), data, meta)
}
