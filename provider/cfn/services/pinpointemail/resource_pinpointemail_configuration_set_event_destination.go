// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourcePinpointEmailConfigurationSetEventDestinationCreate,
		Read:   resourcePinpointEmailConfigurationSetEventDestinationRead,
		Update: resourcePinpointEmailConfigurationSetEventDestinationUpdate,
		Delete: resourcePinpointEmailConfigurationSetEventDestinationDelete,

		Schema: map[string]*schema.Schema{
			"event_destination_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"configuration_set_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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

func resourcePinpointEmailConfigurationSetEventDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::PinpointEmail::ConfigurationSetEventDestination", ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::PinpointEmail::ConfigurationSetEventDestination", ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::PinpointEmail::ConfigurationSetEventDestination", ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::PinpointEmail::ConfigurationSetEventDestination", data, meta)
}