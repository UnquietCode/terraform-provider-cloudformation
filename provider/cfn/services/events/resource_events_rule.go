// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html

package events

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eventsRuleType string = "AWS::Events::Rule"

func ResourceEventsRule() *schema.Resource {
	return &schema.Resource{
		Read: resourceEventsRuleRead,
		Create: resourceEventsRuleCreate,
		Update: resourceEventsRuleUpdate,
		Delete: resourceEventsRuleDelete,
		CustomizeDiff: resourceEventsRuleCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"event_pattern": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"schedule_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"state": {
				Type: schema.TypeString,
				Optional: true,
			},
			"targets": {
				Type: schema.TypeSet,
				Elem: propertyRuleTarget(),
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

func resourceEventsRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eventsRuleType, ResourceEventsRule(), data, meta)
}

func resourceEventsRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eventsRuleType, ResourceEventsRule(), data, meta)
}

func resourceEventsRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eventsRuleType, ResourceEventsRule(), data, meta)
}

func resourceEventsRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eventsRuleType, data, meta)
}

func resourceEventsRuleCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eventsRuleType, data, meta)
}
