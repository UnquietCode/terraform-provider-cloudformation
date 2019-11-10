// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

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
			"input_attachments": {
				Type: schema.TypeList,
				Elem: propertyChannelInputAttachment(),
				Required: false,
			},
			"input_specification": {
				Type: schema.TypeList,
				Elem: propertyChannelInputSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"channel_class": {
				Type: schema.TypeString,
				Required: false,
			},
			"encoder_settings": {
				Type: schema.TypeMap,
				Required: false,
			},
			"destinations": {
				Type: schema.TypeList,
				Elem: propertyChannelOutputDestination(),
				Required: false,
			},
			"log_level": {
				Type: schema.TypeString,
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceMediaLiveChannelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::MediaLive::Channel", data, meta)
}

func resourceMediaLiveChannelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::MediaLive::Channel", data, meta)
}

func resourceMediaLiveChannelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::MediaLive::Channel", data, meta)
}

func resourceMediaLiveChannelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::MediaLive::Channel", data, meta)
}