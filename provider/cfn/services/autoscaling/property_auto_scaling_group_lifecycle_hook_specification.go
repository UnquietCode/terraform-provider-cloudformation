// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAutoScalingGroupLifecycleHookSpecification(extras...string) *schema.Resource {
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
			"default_result": {
				Type: schema.TypeString,
				Required: false,
			},
			"heartbeat_timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
			"lifecycle_hook_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"lifecycle_transition": {
				Type: schema.TypeString,
				Required: true,
			},
			"notification_metadata": {
				Type: schema.TypeString,
				Required: false,
			},
			"notification_target_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}