// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const logsDestinationType string = "AWS::Logs::Destination"

func ResourceLogsDestination() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLogsDestinationExists,
		Read: resourceLogsDestinationRead,
		Create: resourceLogsDestinationCreate,
		Update: resourceLogsDestinationUpdate,
		Delete: resourceLogsDestinationDelete,
		CustomizeDiff: resourceLogsDestinationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"destination_name": {
				Type: schema.TypeString,
				Required: true,
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceLogsDestinationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLogsDestinationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(logsDestinationType, ResourceLogsDestination(), data, meta)
}

func resourceLogsDestinationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(logsDestinationType, ResourceLogsDestination(), data, meta)
}

func resourceLogsDestinationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(logsDestinationType, ResourceLogsDestination(), data, meta)
}

func resourceLogsDestinationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(logsDestinationType, data, meta)
}

func resourceLogsDestinationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(logsDestinationType, data, meta)
}
