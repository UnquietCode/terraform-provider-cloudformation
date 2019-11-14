// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html

package events

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEventsEventBusPolicy() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEventsEventBusPolicyExists,
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
				Type: schema.TypeList,
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
			},
		},
	}
}

func resourceEventsEventBusPolicyExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEventsEventBusPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Events::EventBusPolicy", ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Events::EventBusPolicy", ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Events::EventBusPolicy", ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Events::EventBusPolicy", data, meta)
}

func resourceEventsEventBusPolicyCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}

