// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package iot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceIoTPolicyPrincipalAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceIoTPolicyPrincipalAttachmentCreate,
		Read:   resourceIoTPolicyPrincipalAttachmentRead,
		Update: resourceIoTPolicyPrincipalAttachmentUpdate,
		Delete: resourceIoTPolicyPrincipalAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"policy_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceIoTPolicyPrincipalAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::IoT::PolicyPrincipalAttachment", data, meta)
}

func resourceIoTPolicyPrincipalAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::IoT::PolicyPrincipalAttachment", data, meta)
}

func resourceIoTPolicyPrincipalAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::IoT::PolicyPrincipalAttachment", data, meta)
}

func resourceIoTPolicyPrincipalAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::IoT::PolicyPrincipalAttachment", data, meta)
}