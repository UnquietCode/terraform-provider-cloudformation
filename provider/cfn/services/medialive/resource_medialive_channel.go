// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-channel.html

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMediaLiveChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceMediaLiveChannelCreate,
		Read:   resourceMediaLiveChannelRead,
		Update: resourceMediaLiveChannelUpdate,
		Delete: resourceMediaLiveChannelDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"inputs": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"input_attachments": {
				Type: schema.TypeList,
				Elem: propertyChannelInputAttachment(),
				Optional: true,
			},
			"input_specification": {
				Type: schema.TypeList,
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
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceMediaLiveChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::MediaLive::Channel", ResourceMediaLiveChannel(), data, meta)
}

func resourceMediaLiveChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::MediaLive::Channel", ResourceMediaLiveChannel(), data, meta)
}

func resourceMediaLiveChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::MediaLive::Channel", ResourceMediaLiveChannel(), data, meta)
}

func resourceMediaLiveChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::MediaLive::Channel", data, meta)
}