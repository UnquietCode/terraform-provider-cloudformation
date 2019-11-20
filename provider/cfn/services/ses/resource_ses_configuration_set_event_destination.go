// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-configurationseteventdestination.html

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sESConfigurationSetEventDestinationType string = "AWS::SES::ConfigurationSetEventDestination"

func ResourceSESConfigurationSetEventDestination() *schema.Resource {
	return &schema.Resource{
		Read: resourceSESConfigurationSetEventDestinationRead,
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
				Type: schema.TypeSet,
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

func resourceSESConfigurationSetEventDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sESConfigurationSetEventDestinationType, ResourceSESConfigurationSetEventDestination(), data, meta)
}

func resourceSESConfigurationSetEventDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sESConfigurationSetEventDestinationType, ResourceSESConfigurationSetEventDestination(), data, meta)
}

func resourceSESConfigurationSetEventDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sESConfigurationSetEventDestinationType, ResourceSESConfigurationSetEventDestination(), data, meta)
}

func resourceSESConfigurationSetEventDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sESConfigurationSetEventDestinationType, data, meta)
}

func resourceSESConfigurationSetEventDestinationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sESConfigurationSetEventDestinationType, data, meta)
}
