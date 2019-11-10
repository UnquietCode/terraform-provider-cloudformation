// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTThingPrincipalAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTThingPrincipalAttachmentCreate,
		Read:   resourceIoTThingPrincipalAttachmentRead,
		Delete: resourceIoTThingPrincipalAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"principal": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"thing_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTThingPrincipalAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::ThingPrincipalAttachment", ResourceIoTThingPrincipalAttachment(), data, meta)
}

func resourceIoTThingPrincipalAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::ThingPrincipalAttachment", ResourceIoTThingPrincipalAttachment(), data, meta)
}

func resourceIoTThingPrincipalAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::ThingPrincipalAttachment", ResourceIoTThingPrincipalAttachment(), data, meta)
}

func resourceIoTThingPrincipalAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::ThingPrincipalAttachment", data, meta)
}