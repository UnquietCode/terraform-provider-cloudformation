// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationseteventdestination.html

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSESConfigurationSetEventDestination() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSESConfigurationSetEventDestinationExists,
		Read:   resourceSESConfigurationSetEventDestinationRead,
		Create: resourceSESConfigurationSetEventDestinationCreate,
		Update: resourceSESConfigurationSetEventDestinationUpdate,
		Delete: resourceSESConfigurationSetEventDestinationDelete,
		CustomizeDiff: resourceSESConfigurationSetEventDestinationCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"configuration_set_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"event_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationEventDestination(),
				Required: true,
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

func resourceSESConfigurationSetEventDestinationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSESConfigurationSetEventDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ConfigurationSetEventDestination", ResourceSESConfigurationSetEventDestination(), data, meta)
}

func resourceSESConfigurationSetEventDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ConfigurationSetEventDestination", ResourceSESConfigurationSetEventDestination(), data, meta)
}

func resourceSESConfigurationSetEventDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ConfigurationSetEventDestination", ResourceSESConfigurationSetEventDestination(), data, meta)
}

func resourceSESConfigurationSetEventDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ConfigurationSetEventDestination", data, meta)
}

func resourceSESConfigurationSetEventDestinationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}