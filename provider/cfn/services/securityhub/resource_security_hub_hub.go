// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package securityhub

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSecurityHubHub() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecurityHubHubCreate,
		Read:   resourceSecurityHubHubRead,
		Update: resourceSecurityHubHubUpdate,
		Delete: resourceSecurityHubHubDelete,

		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceSecurityHubHubCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecurityHub::Hub", data, meta)
}

func resourceSecurityHubHubRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecurityHub::Hub", data, meta)
}

func resourceSecurityHubHubUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecurityHub::Hub", data, meta)
}

func resourceSecurityHubHubDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecurityHub::Hub", data, meta)
}