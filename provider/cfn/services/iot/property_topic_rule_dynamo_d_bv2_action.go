// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-dynamodbv2action.html

package iot

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var topicRuleDynamoDBv2ActionProperties map[string]string = map[string]string{
	"put_item": "PutItem",
	"role_arn": "RoleArn",
}

func propertyTopicRuleDynamoDBv2Action(extras...string) *schema.Resource {
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
			"put_item": {
				Type: schema.TypeSet,
				Elem: propertyTopicRulePutItemInput(),
				Optional: true,
				MaxItems: 1,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
