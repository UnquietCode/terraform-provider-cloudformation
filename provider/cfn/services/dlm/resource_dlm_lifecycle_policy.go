// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html

package dlm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const dLMLifecyclePolicyType string = "AWS::DLM::LifecyclePolicy"

var dLMLifecyclePolicyProperties map[string]string = map[string]string{
	"execution_role_arn": "ExecutionRoleArn",
	"description": "Description",
	"state": "State",
	"policy_details": "PolicyDetails",
}

func ResourceDLMLifecyclePolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDLMLifecyclePolicyExists,
		Read: resourceDLMLifecyclePolicyRead,
		Create: resourceDLMLifecyclePolicyCreate,
		Update: resourceDLMLifecyclePolicyUpdate,
		Delete: resourceDLMLifecyclePolicyDelete,
		CustomizeDiff: resourceDLMLifecyclePolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"execution_role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"state": {
				Type: schema.TypeString,
				Optional: true,
			},
			"policy_details": {
				Type: schema.TypeList,
				Elem: propertyLifecyclePolicyPolicyDetails(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDLMLifecyclePolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDLMLifecyclePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dLMLifecyclePolicyType, ResourceDLMLifecyclePolicy(), data, meta)
}

func resourceDLMLifecyclePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dLMLifecyclePolicyType, ResourceDLMLifecyclePolicy(), data, dLMLifecyclePolicyProperties, meta)
}

func resourceDLMLifecyclePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dLMLifecyclePolicyType, ResourceDLMLifecyclePolicy(), data, dLMLifecyclePolicyProperties, meta)
}

func resourceDLMLifecyclePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dLMLifecyclePolicyType, data, meta)
}

func resourceDLMLifecyclePolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dLMLifecyclePolicyType, data, meta)
}
