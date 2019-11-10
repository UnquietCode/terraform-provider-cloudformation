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

func ResourceEIP() *schema.Resource {
	return &schema.Resource{
		Create: resourceEIPCreate,
		Read:   resourceEIPRead,
		Update: resourceEIPUpdate,
		Delete: resourceEIPDelete,

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

func resourceEIPCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::EIP", data, meta)
}

func resourceEIPRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::EIP", data, meta)
}

func resourceEIPUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::EIP", data, meta)
}

func resourceEIPDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::EIP", data, meta)
}