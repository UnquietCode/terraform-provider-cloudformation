// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationseteventdestination.html

package pinpointemail

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourcePinpointEmailConfigurationSetEventDestination() *schema.Resource {
	return &schema.Resource{
		Exists: resourcePinpointEmailConfigurationSetEventDestinationExists,
		Read: resourcePinpointEmailConfigurationSetEventDestinationRead,
		Create: resourcePinpointEmailConfigurationSetEventDestinationCreate,
		Update: resourcePinpointEmailConfigurationSetEventDestinationUpdate,
		Delete: resourcePinpointEmailConfigurationSetEventDestinationDelete,
		CustomizeDiff: resourcePinpointEmailConfigurationSetEventDestinationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"event_destination_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"configuration_set_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"event_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationEventDestination(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourcePinpointEmailConfigurationSetEventDestinationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::PinpointEmail::ConfigurationSetEventDestination", ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::PinpointEmail::ConfigurationSetEventDestination", ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::PinpointEmail::ConfigurationSetEventDestination", ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::PinpointEmail::ConfigurationSetEventDestination", data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

