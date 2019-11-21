// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html

package events

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eventsEventBusPolicyType string = "AWS::Events::EventBusPolicy"

func ResourceEventsEventBusPolicy() *schema.Resource {
	return &schema.Resource{
		Read: resourceEventsEventBusPolicyRead,
		Create: resourceEventsEventBusPolicyCreate,
		Update: resourceEventsEventBusPolicyUpdate,
		Delete: resourceEventsEventBusPolicyDelete,
		CustomizeDiff: resourceEventsEventBusPolicyCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"event_bus_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"condition": {
				Type: schema.TypeSet,
				Elem: propertyEventBusPolicyCondition(),
				Optional: true,
				MaxItems: 1,
			},
			"action": {
				Type: schema.TypeString,
				Required: true,
			},
			"statement_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
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

func resourceEventsEventBusPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eventsEventBusPolicyType, ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eventsEventBusPolicyType, ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eventsEventBusPolicyType, ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eventsEventBusPolicyType, data, meta)
}

func resourceEventsEventBusPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eventsEventBusPolicyType, data, meta)
}
