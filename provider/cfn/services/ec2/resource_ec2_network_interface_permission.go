// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2NetworkInterfacePermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2NetworkInterfacePermissionCreate,
		Read:   resourceEC2NetworkInterfacePermissionRead,
		Update: resourceEC2NetworkInterfacePermissionUpdate,
		Delete: resourceEC2NetworkInterfacePermissionDelete,

		Schema: map[string]*schema.Schema{
			"aws_account_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permission": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2NetworkInterfacePermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NetworkInterfacePermission", data, meta)
}

func resourceEC2NetworkInterfacePermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NetworkInterfacePermission", data, meta)
}

func resourceEC2NetworkInterfacePermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NetworkInterfacePermission", data, meta)
}

func resourceEC2NetworkInterfacePermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NetworkInterfacePermission", data, meta)
}