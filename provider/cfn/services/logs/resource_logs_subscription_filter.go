// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package logs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLogsSubscriptionFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogsSubscriptionFilterCreate,
		Read:   resourceLogsSubscriptionFilterRead,
		Update: resourceLogsSubscriptionFilterUpdate,
		Delete: resourceLogsSubscriptionFilterDelete,

		Schema: map[string]*schema.Schema{
			"destination_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"filter_pattern": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"log_group_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceLogsSubscriptionFilterCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Logs::SubscriptionFilter", data, meta)
}

func resourceLogsSubscriptionFilterRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Logs::SubscriptionFilter", data, meta)
}

func resourceLogsSubscriptionFilterUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Logs::SubscriptionFilter", data, meta)
}

func resourceLogsSubscriptionFilterDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Logs::SubscriptionFilter", data, meta)
}