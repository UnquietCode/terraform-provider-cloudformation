// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html

package sns

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sNSTopicPolicyType string = "AWS::SNS::TopicPolicy"

var sNSTopicPolicyProperties map[string]string = map[string]string{
	"policy_document": "PolicyDocument",
	"topics": "Topics",
}

func ResourceSNSTopicPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSNSTopicPolicyExists,
		Read: resourceSNSTopicPolicyRead,
		Create: resourceSNSTopicPolicyCreate,
		Update: resourceSNSTopicPolicyUpdate,
		Delete: resourceSNSTopicPolicyDelete,
		CustomizeDiff: resourceSNSTopicPolicyCustomizeDiff,
		
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSNSTopicPolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSNSTopicPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sNSTopicPolicyType, ResourceSNSTopicPolicy(), data, meta)
}

func resourceSNSTopicPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sNSTopicPolicyType, ResourceSNSTopicPolicy(), data, sNSTopicPolicyProperties, meta)
}

func resourceSNSTopicPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sNSTopicPolicyType, ResourceSNSTopicPolicy(), data, sNSTopicPolicyProperties, meta)
}

func resourceSNSTopicPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sNSTopicPolicyType, data, meta)
}

func resourceSNSTopicPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sNSTopicPolicyType, data, meta)
}
