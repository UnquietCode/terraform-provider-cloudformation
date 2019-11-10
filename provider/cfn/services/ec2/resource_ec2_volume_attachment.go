// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2VolumeAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2VolumeAttachmentCreate,
		Read:   resourceEC2VolumeAttachmentRead,
		Delete: resourceEC2VolumeAttachmentDelete,

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

func resourceEC2VolumeAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::VolumeAttachment", ResourceEC2VolumeAttachment(), data, meta)
}

func resourceEC2VolumeAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::VolumeAttachment", ResourceEC2VolumeAttachment(), data, meta)
}

func resourceEC2VolumeAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::VolumeAttachment", ResourceEC2VolumeAttachment(), data, meta)
}

func resourceEC2VolumeAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::VolumeAttachment", data, meta)
}