// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iot

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTopicRuleTopicRulePayload(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"actions": {
				Type: schema.TypeSet,
				Elem: propertyTopicRuleAction(),
				Required: true,
			},
			"aws_iot_sql_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"error_action": {
				Type: schema.TypeList,
				Elem: propertyTopicRuleAction(),
				Required: false,
				MaxItems: 1,
			},
			"rule_disabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"sql": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}