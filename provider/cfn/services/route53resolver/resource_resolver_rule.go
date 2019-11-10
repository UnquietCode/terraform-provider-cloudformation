// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package route53resolver

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceResolverRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceResolverRuleCreate,
		Read:   resourceResolverRuleRead,
		Update: resourceResolverRuleUpdate,
		Delete: resourceResolverRuleDelete,

		Schema: map[string]*schema.Schema{
			"resolver_endpoint_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rule_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_ips": {
				Type: schema.TypeList,
				Elem: propertyResolverRuleTargetAddress(),
				Required: false,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceResolverRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53Resolver::ResolverRule", data, meta)
}

func resourceResolverRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53Resolver::ResolverRule", data, meta)
}

func resourceResolverRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53Resolver::ResolverRule", data, meta)
}

func resourceResolverRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53Resolver::ResolverRule", data, meta)
}