// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-input.html

package medialive

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const mediaLiveInputType string = "AWS::MediaLive::Input"

func ResourceMediaLiveInput() *schema.Resource {
	return &schema.Resource{
		Exists: resourceMediaLiveInputExists,
		Read: resourceMediaLiveInputRead,
		Create: resourceMediaLiveInputCreate,
		Update: resourceMediaLiveInputUpdate,
		Delete: resourceMediaLiveInputDelete,
		CustomizeDiff: resourceMediaLiveInputCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"destinations": {
				Type: schema.TypeList,
				Elem: propertyInputInputDestinationRequest(),
				Optional: true,
			},
			"vpc": {
				Type: schema.TypeSet,
				Elem: propertyInputInputVpcRequest(),
				Optional: true,
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
			"sources": {
				Type: schema.TypeList,
				Elem: propertyInputInputSourceRequest(),
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

func resourceMediaLiveInputExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceMediaLiveInputRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(mediaLiveInputType, ResourceMediaLiveInput(), data, meta)
}

func resourceMediaLiveInputCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(mediaLiveInputType, ResourceMediaLiveInput(), data, meta)
}

func resourceMediaLiveInputUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(mediaLiveInputType, ResourceMediaLiveInput(), data, meta)
}

func resourceMediaLiveInputDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(mediaLiveInputType, data, meta)
}

func resourceMediaLiveInputCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(mediaLiveInputType, data, meta)
}
