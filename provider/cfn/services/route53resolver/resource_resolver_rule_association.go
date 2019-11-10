// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package route53resolver

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceResolverRuleAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceResolverRuleAssociationCreate,
		Read:   resourceResolverRuleAssociationRead,
		Update: resourceResolverRuleAssociationUpdate,
		Delete: resourceResolverRuleAssociationDelete,

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resolver_rule_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceResolverRuleAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53Resolver::ResolverRuleAssociation", data, meta)
}

func resourceResolverRuleAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53Resolver::ResolverRuleAssociation", data, meta)
}

func resourceResolverRuleAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53Resolver::ResolverRuleAssociation", data, meta)
}

func resourceResolverRuleAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53Resolver::ResolverRuleAssociation", data, meta)
}