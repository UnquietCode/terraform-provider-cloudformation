// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2EIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2EIPCreate,
		Read:   resourceEC2EIPRead,
		Update: resourceEC2EIPUpdate,
		Delete: resourceEC2EIPDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"public_ipv4_pool": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceEC2EIPCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EIP", data, meta)
}

func resourceEC2EIPRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EIP", data, meta)
}

func resourceEC2EIPUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EIP", data, meta)
}

func resourceEC2EIPDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EIP", data, meta)
}