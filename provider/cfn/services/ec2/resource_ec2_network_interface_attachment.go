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

func ResourceEC2NetworkInterfaceAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2NetworkInterfaceAttachmentCreate,
		Read:   resourceEC2NetworkInterfaceAttachmentRead,
		Update: resourceEC2NetworkInterfaceAttachmentUpdate,
		Delete: resourceEC2NetworkInterfaceAttachmentDelete,

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

func resourceEC2NetworkInterfaceAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::NetworkInterfaceAttachment", data, meta)
}

func resourceEC2NetworkInterfaceAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::NetworkInterfaceAttachment", data, meta)
}

func resourceEC2NetworkInterfaceAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::NetworkInterfaceAttachment", data, meta)
}

func resourceEC2NetworkInterfaceAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::NetworkInterfaceAttachment", data, meta)
}