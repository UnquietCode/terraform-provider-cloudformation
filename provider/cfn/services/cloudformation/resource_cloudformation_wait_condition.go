// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html

package cloudformation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const cloudFormationWaitConditionType string = "AWS::CloudFormation::WaitCondition"

var cloudFormationWaitConditionProperties map[string]string = map[string]string{
	"the_count": "TheCount",
	"handle": "Handle",
	"timeout": "Timeout",
}

func ResourceCloudFormationWaitCondition() *schema.Resource {
	return &schema.Resource{
		Exists: resourceCloudFormationWaitConditionExists,
		Read: resourceCloudFormationWaitConditionRead,
		Create: resourceCloudFormationWaitConditionCreate,
		Update: resourceCloudFormationWaitConditionUpdate,
		Delete: resourceCloudFormationWaitConditionDelete,
		CustomizeDiff: resourceCloudFormationWaitConditionCustomizeDiff,
		
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCloudFormationWaitConditionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceCloudFormationWaitConditionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(cloudFormationWaitConditionType, ResourceCloudFormationWaitCondition(), data, meta)
}

func resourceCloudFormationWaitConditionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(cloudFormationWaitConditionType, ResourceCloudFormationWaitCondition(), data, cloudFormationWaitConditionProperties, meta)
}

func resourceCloudFormationWaitConditionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(cloudFormationWaitConditionType, ResourceCloudFormationWaitCondition(), data, cloudFormationWaitConditionProperties, meta)
}

func resourceCloudFormationWaitConditionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(cloudFormationWaitConditionType, data, meta)
}

func resourceCloudFormationWaitConditionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(cloudFormationWaitConditionType, data, meta)
}
