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

func ResourceWaitConditionHandle() *schema.Resource {
	return &schema.Resource{
		Create: resourceWaitConditionHandleCreate,
		Read:   resourceWaitConditionHandleRead,
		Update: resourceWaitConditionHandleUpdate,
		Delete: resourceWaitConditionHandleDelete,

		Schema: map[string]*schema.Schema{

		},
	}
}

func resourceWaitConditionHandleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::WaitConditionHandle", data, meta)
}

func resourceWaitConditionHandleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::WaitConditionHandle", data, meta)
}

func resourceWaitConditionHandleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::WaitConditionHandle", data, meta)
}

func resourceWaitConditionHandleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::WaitConditionHandle", data, meta)
}