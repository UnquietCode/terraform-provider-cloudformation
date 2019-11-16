// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const mediaLiveChannelType string = "AWS::MediaLive::Channel"

var mediaLiveChannelProperties map[string]string = map[string]string{
	"input_attachments": "InputAttachments",
	"input_specification": "InputSpecification",
	"channel_class": "ChannelClass",
	"encoder_settings": "EncoderSettings",
	"destinations": "Destinations",
	"log_level": "LogLevel",
	"role_arn": "RoleArn",
	"tags": "Tags",
	"name": "Name",
}

func ResourceMediaLiveChannel() *schema.Resource {
	return &schema.Resource{
		Exists: resourceMediaLiveChannelExists,
		Read: resourceMediaLiveChannelRead,
		Create: resourceMediaLiveChannelCreate,
		Update: resourceMediaLiveChannelUpdate,
		Delete: resourceMediaLiveChannelDelete,
		CustomizeDiff: resourceMediaLiveChannelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"input_attachments": {
				Type: schema.TypeList,
				Elem: propertyChannelInputAttachment(),
				Optional: true,
			},
			"input_specification": {
				Type: schema.TypeSet,
				Elem: propertyChannelInputSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"channel_class": {
				Type: schema.TypeString,
				Optional: true,
			},
			"encoder_settings": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"destinations": {
				Type: schema.TypeList,
				Elem: propertyChannelOutputDestination(),
				Optional: true,
			},
			"log_level": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMediaLiveChannelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceMediaLiveChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(mediaLiveChannelType, ResourceMediaLiveChannel(), data, meta)
}

func resourceMediaLiveChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(mediaLiveChannelType, ResourceMediaLiveChannel(), data, mediaLiveChannelProperties, meta)
}

func resourceMediaLiveChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(mediaLiveChannelType, ResourceMediaLiveChannel(), data, mediaLiveChannelProperties, meta)
}

func resourceMediaLiveChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(mediaLiveChannelType, data, meta)
}

func resourceMediaLiveChannelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(mediaLiveChannelType, data, meta)
}
