// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTPolicyType string = "AWS::IoT::Policy"

var ioTPolicyProperties map[string]string = map[string]string{
	"policy_document": "PolicyDocument",
	"policy_name": "PolicyName",
}

func ResourceIoTPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTPolicyExists,
		Read: resourceIoTPolicyRead,
		Create: resourceIoTPolicyCreate,
		Update: resourceIoTPolicyUpdate,
		Delete: resourceIoTPolicyDelete,
		CustomizeDiff: resourceIoTPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"policy_document": {
				Type: schema.TypeMap,
				Required: true,
			},
			"policy_name": {
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

func resourceIoTPolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTPolicyType, ResourceIoTPolicy(), data, meta)
}

func resourceIoTPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTPolicyType, ResourceIoTPolicy(), data, ioTPolicyProperties, meta)
}

func resourceIoTPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTPolicyType, ResourceIoTPolicy(), data, ioTPolicyProperties, meta)
}

func resourceIoTPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTPolicyType, data, meta)
}

func resourceIoTPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTPolicyType, data, meta)
}
