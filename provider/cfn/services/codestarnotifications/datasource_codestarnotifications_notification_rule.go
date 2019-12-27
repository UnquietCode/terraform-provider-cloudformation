// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codestarnotifications-notificationrule.html

package codestarnotifications

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codeStarNotificationsNotificationRuleType string = "AWS::CodeStarNotifications::NotificationRule"

func DatasourceCodeStarNotificationsNotificationRule() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCodeStarNotificationsNotificationRuleRead,
		
		Schema: map[string]*schema.Schema{
			"event_type_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"status": {
				Type: schema.TypeString,
				Optional: true,
			},
			"detail_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource": {
				Type: schema.TypeString,
				Required: true,
			},
			"targets": {
				Type: schema.TypeList,
				Elem: propertyNotificationRuleTarget(),
				Required: true,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCodeStarNotificationsNotificationRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codeStarNotificationsNotificationRuleType, DatasourceCodeStarNotificationsNotificationRule(), data, meta)
}
