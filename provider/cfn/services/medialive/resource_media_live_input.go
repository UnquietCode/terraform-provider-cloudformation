// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

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
			"type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"destinations": {
				Type: schema.TypeList,
				Elem: propertyInputDestinationRequest(),
				Required: false,
			},
			"vpc": {
				Type: schema.TypeList,
				Elem: propertyInputVpcRequest(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"media_connect_flows": {
				Type: schema.TypeList,
				Elem: propertyMediaConnectFlowRequest(),
				Required: false,
			},
			"input_security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"sources": {
				Type: schema.TypeList,
				Elem: propertyInputSourceRequest(),
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

func resourceMediaLiveInputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::MediaLive::Input", data, meta)
}

func resourceMediaLiveInputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::MediaLive::Input", data, meta)
}

func resourceMediaLiveInputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::MediaLive::Input", data, meta)
}

func resourceMediaLiveInputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::MediaLive::Input", data, meta)
}