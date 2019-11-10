// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationset.html

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointEmailConfigurationSet() *schema.Resource {
	return &schema.Resource{
		Create: resourcePinpointEmailConfigurationSetCreate,
		Read:   resourcePinpointEmailConfigurationSetRead,
		Update: resourcePinpointEmailConfigurationSetUpdate,
		Delete: resourcePinpointEmailConfigurationSetDelete,

		Schema: map[string]*schema.Schema{
			"sending_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetSendingOptions(),
				Required: false,
				MaxItems: 1,
			},
			"tracking_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetTrackingOptions(),
				Required: false,
				MaxItems: 1,
			},
			"reputation_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetReputationOptions(),
				Required: false,
				MaxItems: 1,
			},
			"delivery_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetDeliveryOptions(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetTags(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointEmailConfigurationSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::PinpointEmail::ConfigurationSet", data, meta)
}

func resourcePinpointEmailConfigurationSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::PinpointEmail::ConfigurationSet", data, meta)
}

func resourcePinpointEmailConfigurationSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::PinpointEmail::ConfigurationSet", data, meta)
}

func resourcePinpointEmailConfigurationSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::PinpointEmail::ConfigurationSet", data, meta)
}