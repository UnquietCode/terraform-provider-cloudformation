// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html

package route53resolver

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const route53ResolverResolverRuleAssociationType string = "AWS::Route53Resolver::ResolverRuleAssociation"

func ResourceRoute53ResolverResolverRuleAssociation() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceRoute53ResolverResolverRuleAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(route53ResolverResolverRuleAssociationType, ResourceRoute53ResolverResolverRuleAssociation(), data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(route53ResolverResolverRuleAssociationType, ResourceRoute53ResolverResolverRuleAssociation(), data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(route53ResolverResolverRuleAssociationType, ResourceRoute53ResolverResolverRuleAssociation(), data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(route53ResolverResolverRuleAssociationType, data, meta)
}

func resourceRoute53ResolverResolverRuleAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(route53ResolverResolverRuleAssociationType, data, meta)
}
