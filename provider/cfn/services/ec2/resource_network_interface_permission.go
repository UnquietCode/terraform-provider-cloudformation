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

func ResourceNetworkInterfacePermission() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkInterfacePermissionCreate,
		Read:   resourceNetworkInterfacePermissionRead,
		Update: resourceNetworkInterfacePermissionUpdate,
		Delete: resourceNetworkInterfacePermissionDelete,

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

func resourceNetworkInterfacePermissionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NetworkInterfacePermission", data, meta)
}

func resourceNetworkInterfacePermissionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NetworkInterfacePermission", data, meta)
}

func resourceNetworkInterfacePermissionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NetworkInterfacePermission", data, meta)
}

func resourceNetworkInterfacePermissionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NetworkInterfacePermission", data, meta)
}