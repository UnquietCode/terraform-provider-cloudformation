// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2NetworkInterfaceAttachmentType string = "AWS::EC2::NetworkInterfaceAttachment"

var eC2NetworkInterfaceAttachmentProperties map[string]string = map[string]string{
	"delete_on_termination": "DeleteOnTermination",
	"device_index": "DeviceIndex",
	"instance_id": "InstanceId",
	"network_interface_id": "NetworkInterfaceId",
}

func ResourceEC2NetworkInterfaceAttachment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEC2NetworkInterfaceAttachmentExists,
		Read: resourceEC2NetworkInterfaceAttachmentRead,
		Create: resourceEC2NetworkInterfaceAttachmentCreate,
		Update: resourceEC2NetworkInterfaceAttachmentUpdate,
		Delete: resourceEC2NetworkInterfaceAttachmentDelete,
		CustomizeDiff: resourceEC2NetworkInterfaceAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"delete_on_termination": {
				Type: schema.TypeBool,
				Optional: true,
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEC2NetworkInterfaceAttachmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEC2NetworkInterfaceAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2NetworkInterfaceAttachmentType, ResourceEC2NetworkInterfaceAttachment(), data, meta)
}

func resourceEC2NetworkInterfaceAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2NetworkInterfaceAttachmentType, ResourceEC2NetworkInterfaceAttachment(), data, eC2NetworkInterfaceAttachmentProperties, meta)
}

func resourceEC2NetworkInterfaceAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2NetworkInterfaceAttachmentType, ResourceEC2NetworkInterfaceAttachment(), data, eC2NetworkInterfaceAttachmentProperties, meta)
}

func resourceEC2NetworkInterfaceAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2NetworkInterfaceAttachmentType, data, meta)
}

func resourceEC2NetworkInterfaceAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2NetworkInterfaceAttachmentType, data, meta)
}
