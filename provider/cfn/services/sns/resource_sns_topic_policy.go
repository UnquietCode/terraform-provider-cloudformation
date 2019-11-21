// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html

package sns

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sNSTopicPolicyType string = "AWS::SNS::TopicPolicy"

func ResourceSNSTopicPolicy() *schema.Resource {
	return &schema.Resource{
		Read: resourceSNSTopicPolicyRead,
		Create: resourceSNSTopicPolicyCreate,
		Update: resourceSNSTopicPolicyUpdate,
		Delete: resourceSNSTopicPolicyDelete,
		CustomizeDiff: resourceSNSTopicPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"policy_document": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceSNSTopicPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sNSTopicPolicyType, ResourceSNSTopicPolicy(), data, meta)
}

func resourceSNSTopicPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sNSTopicPolicyType, ResourceSNSTopicPolicy(), data, meta)
}

func resourceSNSTopicPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sNSTopicPolicyType, ResourceSNSTopicPolicy(), data, meta)
}

func resourceSNSTopicPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sNSTopicPolicyType, data, meta)
}

func resourceSNSTopicPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sNSTopicPolicyType, data, meta)
}
