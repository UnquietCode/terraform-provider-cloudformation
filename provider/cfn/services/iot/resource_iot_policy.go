// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html

package iot

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTPolicyType string = "AWS::IoT::Policy"

func ResourceIoTPolicy() *schema.Resource {
	return &schema.Resource{
		Read: resourceIoTPolicyRead,
		Create: resourceIoTPolicyCreate,
		Update: resourceIoTPolicyUpdate,
		Delete: resourceIoTPolicyDelete,
		CustomizeDiff: resourceIoTPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"policy_document": {
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"policy_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceIoTPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTPolicyType, ResourceIoTPolicy(), data, meta)
}

func resourceIoTPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTPolicyType, ResourceIoTPolicy(), data, meta)
}

func resourceIoTPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTPolicyType, ResourceIoTPolicy(), data, meta)
}

func resourceIoTPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTPolicyType, data, meta)
}

func resourceIoTPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTPolicyType, data, meta)
}
