// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceMediaLiveInput() *schema.Resource {
	return &schema.Resource{
		Create: resourceMediaLiveInputCreate,
		Read:   resourceMediaLiveInputRead,
		Update: resourceMediaLiveInputUpdate,
		Delete: resourceMediaLiveInputDelete,

		Schema: map[string]*schema.Schema{
			"destinations": {
				Type: schema.TypeList,
				Elem: propertyInputInputDestinationRequest(),
				Optional: true,
				Computed: true,
			},
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertyInputInputSourceRequest(),
				Optional: true,
				Computed: true,
			},
			"type": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vpc": {
				Type: schema.TypeList,
				Elem: propertyInputInputVpcRequest(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"media_connect_flows": {
				Type: schema.TypeList,
				Elem: propertyInputMediaConnectFlowRequest(),
				Optional: true,
			},
			"input_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceMediaLiveInputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::MediaLive::Input", ResourceMediaLiveInput(), data, meta)
}

func resourceMediaLiveInputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::MediaLive::Input", ResourceMediaLiveInput(), data, meta)
}

func resourceMediaLiveInputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::MediaLive::Input", ResourceMediaLiveInput(), data, meta)
}

func resourceMediaLiveInputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::MediaLive::Input", data, meta)
}