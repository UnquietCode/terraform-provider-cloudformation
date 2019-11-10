// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLogStream() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogStreamCreate,
		Read:   resourceLogStreamRead,
		Update: resourceLogStreamUpdate,
		Delete: resourceLogStreamDelete,

		Schema: map[string]*schema.Schema{
			"log_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"log_stream_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceLogStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::LogStream", data, meta)
}

func resourceLogStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::LogStream", data, meta)
}

func resourceLogStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::LogStream", data, meta)
}

func resourceLogStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::LogStream", data, meta)
}