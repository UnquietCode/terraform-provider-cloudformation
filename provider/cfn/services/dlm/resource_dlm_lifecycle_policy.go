// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dlm-lifecyclepolicy.html

package dlm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDLMLifecyclePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDLMLifecyclePolicyCreate,
		Read:   resourceDLMLifecyclePolicyRead,
		Update: resourceDLMLifecyclePolicyUpdate,
		Delete: resourceDLMLifecyclePolicyDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
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
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDLMLifecyclePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DLM::LifecyclePolicy", ResourceDLMLifecyclePolicy(), data, meta)
}

func resourceDLMLifecyclePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DLM::LifecyclePolicy", ResourceDLMLifecyclePolicy(), data, meta)
}

func resourceDLMLifecyclePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DLM::LifecyclePolicy", ResourceDLMLifecyclePolicy(), data, meta)
}

func resourceDLMLifecyclePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DLM::LifecyclePolicy", data, meta)
}