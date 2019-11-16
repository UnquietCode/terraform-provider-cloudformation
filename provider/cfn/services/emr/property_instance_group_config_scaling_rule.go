// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html

package emr

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var instanceGroupConfigScalingRuleProperties map[string]string = map[string]string{
	"action": "Action",
	"description": "Description",
	"name": "Name",
	"trigger": "Trigger",
}

func propertyInstanceGroupConfigScalingRule(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action": {
				Type: schema.TypeList,
				Elem: propertyInstanceGroupConfigScalingAction(),
				Required: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"trigger": {
				Type: schema.TypeList,
				Elem: propertyInstanceGroupConfigScalingTrigger(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}
