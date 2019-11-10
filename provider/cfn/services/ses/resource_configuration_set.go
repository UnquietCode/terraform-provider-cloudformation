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

func ResourceConfigurationSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationSetCreate,
		Read:   resourceConfigurationSetRead,
		Update: resourceConfigurationSetUpdate,
		Delete: resourceConfigurationSetDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceConfigurationSetCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SES::ConfigurationSet", data, meta)
}

func resourceConfigurationSetRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SES::ConfigurationSet", data, meta)
}

func resourceConfigurationSetUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SES::ConfigurationSet", data, meta)
}

func resourceConfigurationSetDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SES::ConfigurationSet", data, meta)
}