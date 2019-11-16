// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const logsSubscriptionFilterType string = "AWS::Logs::SubscriptionFilter"

var logsSubscriptionFilterProperties map[string]string = map[string]string{
	"destination_arn": "DestinationArn",
	"filter_pattern": "FilterPattern",
	"log_group_name": "LogGroupName",
	"role_arn": "RoleArn",
}

func ResourceLogsSubscriptionFilter() *schema.Resource {
	return &schema.Resource{
		Exists: resourceLogsSubscriptionFilterExists,
		Read: resourceLogsSubscriptionFilterRead,
		Create: resourceLogsSubscriptionFilterCreate,
		Update: resourceLogsSubscriptionFilterUpdate,
		Delete: resourceLogsSubscriptionFilterDelete,
		CustomizeDiff: resourceLogsSubscriptionFilterCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"destination_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"filter_pattern": {
				Type: schema.TypeString,
				Required: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"role_arn": {
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

func resourceLogsSubscriptionFilterExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceLogsSubscriptionFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(logsSubscriptionFilterType, ResourceLogsSubscriptionFilter(), data, meta)
}

func resourceLogsSubscriptionFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(logsSubscriptionFilterType, ResourceLogsSubscriptionFilter(), data, logsSubscriptionFilterProperties, meta)
}

func resourceLogsSubscriptionFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(logsSubscriptionFilterType, ResourceLogsSubscriptionFilter(), data, logsSubscriptionFilterProperties, meta)
}

func resourceLogsSubscriptionFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(logsSubscriptionFilterType, data, meta)
}

func resourceLogsSubscriptionFilterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(logsSubscriptionFilterType, data, meta)
}
