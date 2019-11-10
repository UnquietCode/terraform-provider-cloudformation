// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLogsDestination() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogsDestinationCreate,
		Read:   resourceLogsDestinationRead,
		Update: resourceLogsDestinationUpdate,
		Delete: resourceLogsDestinationDelete,

		Schema: map[string]*schema.Schema{
			"arn": {
				Type: schema.TypeString,
				Computed: true,
			},
			"destination_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"destination_policy": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceLogsDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::Destination", ResourceLogsDestination(), data, meta)
}

func resourceLogsDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::Destination", ResourceLogsDestination(), data, meta)
}

func resourceLogsDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::Destination", ResourceLogsDestination(), data, meta)
}

func resourceLogsDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::Destination", data, meta)
}