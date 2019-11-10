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

func ResourceThingPrincipalAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceThingPrincipalAttachmentCreate,
		Read:   resourceThingPrincipalAttachmentRead,
		Update: resourceThingPrincipalAttachmentUpdate,
		Delete: resourceThingPrincipalAttachmentDelete,

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

func resourceThingPrincipalAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::ThingPrincipalAttachment", data, meta)
}

func resourceThingPrincipalAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::ThingPrincipalAttachment", data, meta)
}

func resourceThingPrincipalAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::ThingPrincipalAttachment", data, meta)
}

func resourceThingPrincipalAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::ThingPrincipalAttachment", data, meta)
}