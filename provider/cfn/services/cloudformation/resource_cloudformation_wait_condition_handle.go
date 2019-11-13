// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCloudFormationWaitConditionHandle() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCloudFormationWaitConditionHandleExists,
		Read:   resourceCloudFormationWaitConditionHandleRead,
		Create: resourceCloudFormationWaitConditionHandleCreate,
		Delete: resourceCloudFormationWaitConditionHandleDelete,
		
		Schema: map[string]*schema.Schema{
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudFormationWaitConditionHandleExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCloudFormationWaitConditionHandleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CloudFormation::WaitConditionHandle", ResourceCloudFormationWaitConditionHandle(), data, meta)
}

func resourceCloudFormationWaitConditionHandleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CloudFormation::WaitConditionHandle", ResourceCloudFormationWaitConditionHandle(), data, meta)
}

func resourceCloudFormationWaitConditionHandleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CloudFormation::WaitConditionHandle", ResourceCloudFormationWaitConditionHandle(), data, meta)
}

func resourceCloudFormationWaitConditionHandleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CloudFormation::WaitConditionHandle", data, meta)
}