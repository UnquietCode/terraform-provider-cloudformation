// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2EIP() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2EIPExists,
		Read:   resourceEC2EIPRead,
		Create: resourceEC2EIPCreate,
		Update: resourceEC2EIPUpdate,
		Delete: resourceEC2EIPDelete,
		CustomizeDiff: resourceEC2EIPCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"public_ipv4_pool": {
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

func resourceEC2EIPExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2EIPRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EIP", ResourceEC2EIP(), data, meta)
}

func resourceEC2EIPCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EIP", ResourceEC2EIP(), data, meta)
}

func resourceEC2EIPUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EIP", ResourceEC2EIP(), data, meta)
}

func resourceEC2EIPDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EIP", data, meta)
}

func resourceEC2EIPCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}