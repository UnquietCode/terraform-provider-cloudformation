// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package securityhub

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceHub() *schema.Resource {
	return &schema.Resource{
		Create: resourceHubCreate,
		Read:   resourceHubRead,
		Update: resourceHubUpdate,
		Delete: resourceHubDelete,

		Schema: map[string]*schema.Schema{
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceHubCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SecurityHub::Hub", data, meta)
}

func resourceHubRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SecurityHub::Hub", data, meta)
}

func resourceHubUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SecurityHub::Hub", data, meta)
}

func resourceHubDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SecurityHub::Hub", data, meta)
}