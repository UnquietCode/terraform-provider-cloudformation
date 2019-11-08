// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSESConfigurationSetEventDestination() *schema.Resource {
	return &schema.Resource{
		Create: resourceSESConfigurationSetEventDestinationCreate,
		Read:   resourceSESConfigurationSetEventDestinationRead,
		Update: resourceSESConfigurationSetEventDestinationUpdate,
		Delete: resourceSESConfigurationSetEventDestinationDelete,

		Schema: map[string]*schema.Schema{
			"configuration_set_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"event_destination": {
				Type: schema.TypeList,
				Elem: propertyEventDestination(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceSESConfigurationSetEventDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ConfigurationSetEventDestination", data, meta)
}

func resourceSESConfigurationSetEventDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ConfigurationSetEventDestination", data, meta)
}

func resourceSESConfigurationSetEventDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ConfigurationSetEventDestination", data, meta)
}

func resourceSESConfigurationSetEventDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ConfigurationSetEventDestination", data, meta)
}