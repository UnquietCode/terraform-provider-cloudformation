// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package wafregional

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyWebACLRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeList,
				Elem: propertyWebACLAction(),
				Required: true,
				MaxItems: 1,
			},
			"priority": {
				Type: schema.TypeInt,
				Required: true,
			},
			"rule_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}