// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTThingPrincipalAttachmentType string = "AWS::IoT::ThingPrincipalAttachment"

var ioTThingPrincipalAttachmentProperties map[string]string = map[string]string{
	"principal": "Principal",
	"thing_name": "ThingName",
}

func ResourceIoTThingPrincipalAttachment() *schema.Resource {
	return &schema.Resource{
		Exists: resourceIoTThingPrincipalAttachmentExists,
		Read: resourceIoTThingPrincipalAttachmentRead,
		Create: resourceIoTThingPrincipalAttachmentCreate,
		Update: resourceIoTThingPrincipalAttachmentUpdate,
		Delete: resourceIoTThingPrincipalAttachmentDelete,
		CustomizeDiff: resourceIoTThingPrincipalAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"principal": {
				Type: schema.TypeString,
				Required: true,
			},
			"thing_name": {
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

func resourceIoTThingPrincipalAttachmentExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceIoTThingPrincipalAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTThingPrincipalAttachmentType, ResourceIoTThingPrincipalAttachment(), data, meta)
}

func resourceIoTThingPrincipalAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTThingPrincipalAttachmentType, ResourceIoTThingPrincipalAttachment(), data, ioTThingPrincipalAttachmentProperties, meta)
}

func resourceIoTThingPrincipalAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTThingPrincipalAttachmentType, ResourceIoTThingPrincipalAttachment(), data, ioTThingPrincipalAttachmentProperties, meta)
}

func resourceIoTThingPrincipalAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTThingPrincipalAttachmentType, data, meta)
}

func resourceIoTThingPrincipalAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTThingPrincipalAttachmentType, data, meta)
}
