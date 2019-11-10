// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package events

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRuleEcsParameters(extras...string) *schema.Resource {
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
			"group": {
				Type: schema.TypeString,
				Required: false,
			},
			"launch_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"network_configuration": {
				Type: schema.TypeList,
				Elem: propertyRuleNetworkConfiguration(),
				Required: false,
				MaxItems: 1,
			},
			"platform_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"task_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"task_definition_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}