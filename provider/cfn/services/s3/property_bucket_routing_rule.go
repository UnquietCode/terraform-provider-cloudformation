// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketRoutingRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"redirect_rule": {
				Type: schema.TypeList,
				Elem: propertyBucketRedirectRule(),
				Required: true,
				MaxItems: 1,
			},
			"routing_rule_condition": {
				Type: schema.TypeList,
				Elem: propertyBucketRoutingRuleCondition(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}