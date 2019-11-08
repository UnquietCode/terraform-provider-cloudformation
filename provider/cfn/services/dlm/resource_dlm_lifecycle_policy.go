// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

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
			"execution_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"state": {
				Type: schema.TypeString,
				Required: false,
			},
			"policy_details": {
				Type: schema.TypeList,
				Elem: propertyPolicyDetails(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceDLMLifecyclePolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DLM::LifecyclePolicy", data, meta)
}

func resourceDLMLifecyclePolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DLM::LifecyclePolicy", data, meta)
}

func resourceDLMLifecyclePolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DLM::LifecyclePolicy", data, meta)
}

func resourceDLMLifecyclePolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DLM::LifecyclePolicy", data, meta)
}