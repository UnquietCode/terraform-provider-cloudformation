// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html

package logs

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const logsSubscriptionFilterType string = "AWS::Logs::SubscriptionFilter"

func ResourceLogsSubscriptionFilter() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceLogsSubscriptionFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(logsSubscriptionFilterType, ResourceLogsSubscriptionFilter(), data, meta)
}

func resourceLogsSubscriptionFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(logsSubscriptionFilterType, ResourceLogsSubscriptionFilter(), data, meta)
}

func resourceLogsSubscriptionFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(logsSubscriptionFilterType, ResourceLogsSubscriptionFilter(), data, meta)
}

func resourceLogsSubscriptionFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(logsSubscriptionFilterType, data, meta)
}

func resourceLogsSubscriptionFilterCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(logsSubscriptionFilterType, data, meta)
}
