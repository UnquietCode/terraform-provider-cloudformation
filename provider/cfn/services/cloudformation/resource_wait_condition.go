// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceWaitCondition() *schema.Resource {
	return &schema.Resource{
		Create: resourceWaitConditionCreate,
		Read:   resourceWaitConditionRead,
		Update: resourceWaitConditionUpdate,
		Delete: resourceWaitConditionDelete,

		Schema: map[string]*schema.Schema{
			"count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"handle": {
				Type: schema.TypeString,
				Required: false,
			},
			"timeout": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceWaitConditionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::WaitCondition", data, meta)
}

func resourceWaitConditionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::WaitCondition", data, meta)
}

func resourceWaitConditionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::WaitCondition", data, meta)
}

func resourceWaitConditionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::WaitCondition", data, meta)
}