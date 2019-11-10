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

func ResourceVolumeAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolumeAttachmentCreate,
		Read:   resourceVolumeAttachmentRead,
		Update: resourceVolumeAttachmentUpdate,
		Delete: resourceVolumeAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"device": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"volume_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceVolumeAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VolumeAttachment", data, meta)
}

func resourceVolumeAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VolumeAttachment", data, meta)
}

func resourceVolumeAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VolumeAttachment", data, meta)
}

func resourceVolumeAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VolumeAttachment", data, meta)
}