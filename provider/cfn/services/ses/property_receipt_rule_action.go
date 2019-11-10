// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyReceiptRuleAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bounce_action": {
				Type: schema.TypeList,
				Elem: propertyReceiptRuleBounceAction(),
				Required: false,
				MaxItems: 1,
			},
			"s3_action": {
				Type: schema.TypeList,
				Elem: propertyReceiptRuleS3Action(),
				Required: false,
				MaxItems: 1,
			},
			"stop_action": {
				Type: schema.TypeList,
				Elem: propertyReceiptRuleStopAction(),
				Required: false,
				MaxItems: 1,
			},
			"sns_action": {
				Type: schema.TypeList,
				Elem: propertyReceiptRuleSNSAction(),
				Required: false,
				MaxItems: 1,
			},
			"workmail_action": {
				Type: schema.TypeList,
				Elem: propertyReceiptRuleWorkmailAction(),
				Required: false,
				MaxItems: 1,
			},
			"add_header_action": {
				Type: schema.TypeList,
				Elem: propertyReceiptRuleAddHeaderAction(),
				Required: false,
				MaxItems: 1,
			},
			"lambda_action": {
				Type: schema.TypeList,
				Elem: propertyReceiptRuleLambdaAction(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}