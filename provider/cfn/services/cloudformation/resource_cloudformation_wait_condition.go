// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudFormationWaitCondition() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFormationWaitConditionCreate,
		Read:   resourceCloudFormationWaitConditionRead,
		Update: resourceCloudFormationWaitConditionUpdate,
		Delete: resourceCloudFormationWaitConditionDelete,

		Schema: map[string]*schema.Schema{
			"the_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"handle": {
				Type: schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCloudFormationWaitConditionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::WaitCondition", data, meta)
}

func resourceCloudFormationWaitConditionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::WaitCondition", data, meta)
}

func resourceCloudFormationWaitConditionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::WaitCondition", data, meta)
}

func resourceCloudFormationWaitConditionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::WaitCondition", data, meta)
}