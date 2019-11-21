// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html

package iot

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const ioTPolicyPrincipalAttachmentType string = "AWS::IoT::PolicyPrincipalAttachment"

func ResourceIoTPolicyPrincipalAttachment() *schema.Resource {
	return &schema.Resource{
		Read: resourceIoTPolicyPrincipalAttachmentRead,
		Create: resourceIoTPolicyPrincipalAttachmentCreate,
		Update: resourceIoTPolicyPrincipalAttachmentUpdate,
		Delete: resourceIoTPolicyPrincipalAttachmentDelete,
		CustomizeDiff: resourceIoTPolicyPrincipalAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"policy_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceIoTPolicyPrincipalAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(ioTPolicyPrincipalAttachmentType, ResourceIoTPolicyPrincipalAttachment(), data, meta)
}

func resourceIoTPolicyPrincipalAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(ioTPolicyPrincipalAttachmentType, ResourceIoTPolicyPrincipalAttachment(), data, meta)
}

func resourceIoTPolicyPrincipalAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(ioTPolicyPrincipalAttachmentType, ResourceIoTPolicyPrincipalAttachment(), data, meta)
}

func resourceIoTPolicyPrincipalAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(ioTPolicyPrincipalAttachmentType, data, meta)
}

func resourceIoTPolicyPrincipalAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(ioTPolicyPrincipalAttachmentType, data, meta)
}
