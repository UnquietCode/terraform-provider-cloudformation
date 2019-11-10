// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLogsLogGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogsLogGroupCreate,
		Read:   resourceLogsLogGroupRead,
		Update: resourceLogsLogGroupUpdate,
		Delete: resourceLogsLogGroupDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"retention_in_days": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceLogsLogGroupCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::LogGroup", data, meta)
}

func resourceLogsLogGroupRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::LogGroup", data, meta)
}

func resourceLogsLogGroupUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::LogGroup", data, meta)
}

func resourceLogsLogGroupDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::LogGroup", data, meta)
}