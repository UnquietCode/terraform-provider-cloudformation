// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sns

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceTopicPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceTopicPolicyCreate,
		Read:   resourceTopicPolicyRead,
		Update: resourceTopicPolicyUpdate,
		Delete: resourceTopicPolicyDelete,

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

func resourceTopicPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SNS::TopicPolicy", data, meta)
}

func resourceTopicPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SNS::TopicPolicy", data, meta)
}

func resourceTopicPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SNS::TopicPolicy", data, meta)
}

func resourceTopicPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SNS::TopicPolicy", data, meta)
}