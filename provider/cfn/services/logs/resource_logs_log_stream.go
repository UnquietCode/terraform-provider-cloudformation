// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const logsLogStreamType string = "AWS::Logs::LogStream"

var logsLogStreamProperties map[string]string = map[string]string{
	"log_group_name": "LogGroupName",
	"log_stream_name": "LogStreamName",
}

func ResourceLogsLogStream() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLogsLogStreamExists,
		Read: resourceLogsLogStreamRead,
		Create: resourceLogsLogStreamCreate,
		Update: resourceLogsLogStreamUpdate,
		Delete: resourceLogsLogStreamDelete,
		CustomizeDiff: resourceLogsLogStreamCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"log_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"log_stream_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLogsLogStreamExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLogsLogStreamRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(logsLogStreamType, ResourceLogsLogStream(), data, meta)
}

func resourceLogsLogStreamCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(logsLogStreamType, ResourceLogsLogStream(), data, logsLogStreamProperties, meta)
}

func resourceLogsLogStreamUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(logsLogStreamType, ResourceLogsLogStream(), data, logsLogStreamProperties, meta)
}

func resourceLogsLogStreamDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(logsLogStreamType, data, meta)
}

func resourceLogsLogStreamCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(logsLogStreamType, data, meta)
}
