// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
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
		Create: resourceEventsEventBusPolicyCreate,
		Read:   resourceEventsEventBusPolicyRead,
		Update: resourceEventsEventBusPolicyUpdate,
		Delete: resourceEventsEventBusPolicyDelete,

		Schema: map[string]*schema.Schema{
			"event_bus_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"principal": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEventsEventBusPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Events::EventBusPolicy", ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Events::EventBusPolicy", ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Events::EventBusPolicy", ResourceEventsEventBusPolicy(), data, meta)
}

func resourceEventsEventBusPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Events::EventBusPolicy", data, meta)
}