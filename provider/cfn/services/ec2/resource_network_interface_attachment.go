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

func ResourceNetworkInterfaceAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkInterfaceAttachmentCreate,
		Read:   resourceNetworkInterfaceAttachmentRead,
		Update: resourceNetworkInterfaceAttachmentUpdate,
		Delete: resourceNetworkInterfaceAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"delete_on_termination": {
				Type: schema.TypeBool,
				Required: false,
			},
			"device_index": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"network_interface_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNetworkInterfaceAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NetworkInterfaceAttachment", data, meta)
}

func resourceNetworkInterfaceAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NetworkInterfaceAttachment", data, meta)
}

func resourceNetworkInterfaceAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NetworkInterfaceAttachment", data, meta)
}

func resourceNetworkInterfaceAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NetworkInterfaceAttachment", data, meta)
}