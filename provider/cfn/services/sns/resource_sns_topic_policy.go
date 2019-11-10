// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html

package sns

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSNSTopicPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSNSTopicPolicyCreate,
		Read:   resourceSNSTopicPolicyRead,
		Update: resourceSNSTopicPolicyUpdate,
		Delete: resourceSNSTopicPolicyDelete,

		Schema: map[string]*schema.Schema{
			"policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
			"topics": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
		},
	}
}

func resourceSNSTopicPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SNS::TopicPolicy", ResourceSNSTopicPolicy(), data, meta)
}

func resourceSNSTopicPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SNS::TopicPolicy", ResourceSNSTopicPolicy(), data, meta)
}

func resourceSNSTopicPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SNS::TopicPolicy", ResourceSNSTopicPolicy(), data, meta)
}

func resourceSNSTopicPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SNS::TopicPolicy", data, meta)
}