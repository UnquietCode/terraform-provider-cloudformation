// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2EIPType string = "AWS::EC2::EIP"

func ResourceEC2EIP() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2EIPRead,
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

func resourceEC2EIPRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2EIPType, ResourceEC2EIP(), data, meta)
}

func resourceEC2EIPCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2EIPType, ResourceEC2EIP(), data, meta)
}

func resourceEC2EIPUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2EIPType, ResourceEC2EIP(), data, meta)
}

func resourceEC2EIPDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2EIPType, data, meta)
}

func resourceEC2EIPCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2EIPType, data, meta)
}
