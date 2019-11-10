// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigurationSetEventDestination() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationSetEventDestinationCreate,
		Read:   resourceConfigurationSetEventDestinationRead,
		Update: resourceConfigurationSetEventDestinationUpdate,
		Delete: resourceConfigurationSetEventDestinationDelete,

		Schema: map[string]*schema.Schema{
			"configuration_set_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"event_destination": {
				Type: schema.TypeList,
				Elem: propertyConfigurationSetEventDestinationEventDestination(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceConfigurationSetEventDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ConfigurationSetEventDestination", data, meta)
}

func resourceConfigurationSetEventDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ConfigurationSetEventDestination", data, meta)
}

func resourceConfigurationSetEventDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ConfigurationSetEventDestination", data, meta)
}

func resourceConfigurationSetEventDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ConfigurationSetEventDestination", data, meta)
}