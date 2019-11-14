// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html

package route53resolver

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRoute53ResolverResolverRuleAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRoute53ResolverResolverRuleAssociationExists,
		Read: resourceRoute53ResolverResolverRuleAssociationRead,
		Create: resourceRoute53ResolverResolverRuleAssociationCreate,
		Update: resourceRoute53ResolverResolverRuleAssociationUpdate,
		Delete: resourceRoute53ResolverResolverRuleAssociationDelete,
		CustomizeDiff: resourceRoute53ResolverResolverRuleAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"resolver_rule_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
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

func resourceRoute53ResolverResolverRuleAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Route53Resolver::ResolverRuleAssociation", ResourceRoute53ResolverResolverRuleAssociation(), data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Route53Resolver::ResolverRuleAssociation", ResourceRoute53ResolverResolverRuleAssociation(), data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Route53Resolver::ResolverRuleAssociation", ResourceRoute53ResolverResolverRuleAssociation(), data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Route53Resolver::ResolverRuleAssociation", data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

