// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-ecsparameters.html

package events

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var ruleEcsParametersProperties map[string]string = map[string]string{
	"group": "Group",
	"launch_type": "LaunchType",
	"network_configuration": "NetworkConfiguration",
	"platform_version": "PlatformVersion",
	"task_count": "TaskCount",
	"task_definition_arn": "TaskDefinitionArn",
}

func propertyRuleEcsParameters(extras...string) *schema.Resource {
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
			"group": {
				Type: schema.TypeString,
				Optional: true,
			},
			"launch_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"network_configuration": {
				Type: schema.TypeList,
				Elem: propertyRuleNetworkConfiguration(),
				Optional: true,
				MaxItems: 1,
			},
			"platform_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"task_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"task_definition_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}
