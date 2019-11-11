// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLogsLogStream() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogsLogStreamCreate,
		Read:   resourceLogsLogStreamRead,
		Delete: resourceLogsLogStreamDelete,

		Schema: map[string]*schema.Schema{
			"log_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"log_stream_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLogsLogStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::LogStream", ResourceLogsLogStream(), data, meta)
}

func resourceLogsLogStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::LogStream", ResourceLogsLogStream(), data, meta)
}

func resourceLogsLogStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::LogStream", ResourceLogsLogStream(), data, meta)
}

func resourceLogsLogStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::LogStream", data, meta)
}