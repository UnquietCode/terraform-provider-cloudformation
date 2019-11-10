// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTThingPrincipalAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTThingPrincipalAttachmentCreate,
		Read:   resourceIoTThingPrincipalAttachmentRead,
		Update: resourceIoTThingPrincipalAttachmentUpdate,
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
	return plugin.ResourceCreate("AWS::IoT::ThingPrincipalAttachment", data, meta)
}

func resourceIoTThingPrincipalAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::ThingPrincipalAttachment", data, meta)
}

func resourceIoTThingPrincipalAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::ThingPrincipalAttachment", data, meta)
}

func resourceIoTThingPrincipalAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::ThingPrincipalAttachment", data, meta)
}