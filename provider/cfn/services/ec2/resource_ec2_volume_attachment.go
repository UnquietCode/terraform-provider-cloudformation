// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2VolumeAttachmentType string = "AWS::EC2::VolumeAttachment"

func ResourceEC2VolumeAttachment() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2VolumeAttachmentRead,
		Create: resourceEC2VolumeAttachmentCreate,
		Update: resourceEC2VolumeAttachmentUpdate,
		Delete: resourceEC2VolumeAttachmentDelete,
		CustomizeDiff: resourceEC2VolumeAttachmentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"device": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"volume_id": {
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

func resourceEC2VolumeAttachmentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2VolumeAttachmentType, ResourceEC2VolumeAttachment(), data, meta)
}

func resourceEC2VolumeAttachmentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2VolumeAttachmentType, ResourceEC2VolumeAttachment(), data, meta)
}

func resourceEC2VolumeAttachmentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2VolumeAttachmentType, ResourceEC2VolumeAttachment(), data, meta)
}

func resourceEC2VolumeAttachmentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2VolumeAttachmentType, data, meta)
}

func resourceEC2VolumeAttachmentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2VolumeAttachmentType, data, meta)
}
