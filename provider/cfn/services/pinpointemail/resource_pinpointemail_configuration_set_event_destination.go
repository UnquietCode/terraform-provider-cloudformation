// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-configurationseteventdestination.html

package pinpointemail

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const pinpointEmailConfigurationSetEventDestinationType string = "AWS::PinpointEmail::ConfigurationSetEventDestination"

func ResourcePinpointEmailConfigurationSetEventDestination() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeSet,
				Elem: propertyConfigurationSetEventDestinationEventDestination(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourcePinpointEmailConfigurationSetEventDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(pinpointEmailConfigurationSetEventDestinationType, ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(pinpointEmailConfigurationSetEventDestinationType, ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(pinpointEmailConfigurationSetEventDestinationType, ResourcePinpointEmailConfigurationSetEventDestination(), data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(pinpointEmailConfigurationSetEventDestinationType, data, meta)
}

func resourcePinpointEmailConfigurationSetEventDestinationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(pinpointEmailConfigurationSetEventDestinationType, data, meta)
}
