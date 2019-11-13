// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourcePinpointEmailConfigurationSetExists,
		Read:   resourcePinpointEmailConfigurationSetRead,
		Create: resourcePinpointEmailConfigurationSetCreate,
		Update: resourcePinpointEmailConfigurationSetUpdate,
		Delete: resourcePinpointEmailConfigurationSetDelete,
		
		Schema: map[string]*schema.Schema{
			"sending_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetSendingOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"tracking_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetTrackingOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"reputation_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetReputationOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"delivery_options": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetDeliveryOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetTags(),
				Optional: true,
			},
			"name": {
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

func resourcePinpointEmailConfigurationSetExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointEmailConfigurationSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::PinpointEmail::ConfigurationSet", ResourcePinpointEmailConfigurationSet(), data, meta)
}

func resourcePinpointEmailConfigurationSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::PinpointEmail::ConfigurationSet", ResourcePinpointEmailConfigurationSet(), data, meta)
}

func resourcePinpointEmailConfigurationSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::PinpointEmail::ConfigurationSet", ResourcePinpointEmailConfigurationSet(), data, meta)
}

func resourcePinpointEmailConfigurationSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::PinpointEmail::ConfigurationSet", data, meta)
}