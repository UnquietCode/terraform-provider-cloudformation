// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html

package events

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEventsEventBus() *schema.Resource {
	return &schema.Resource{
		Exists: resourceEventsEventBusExists,
		Read: resourceEventsEventBusRead,
		Create: resourceEventsEventBusCreate,
		Update: resourceEventsEventBusUpdate,
		Delete: resourceEventsEventBusDelete,
		CustomizeDiff: resourceEventsEventBusCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"event_source_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
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

func resourceEventsEventBusExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceEventsEventBusRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Events::EventBus", ResourceEventsEventBus(), data, meta)
}

func resourceEventsEventBusCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Events::EventBus", ResourceEventsEventBus(), data, meta)
}

func resourceEventsEventBusUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Events::EventBus", ResourceEventsEventBus(), data, meta)
}

func resourceEventsEventBusDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Events::EventBus", data, meta)
}

func resourceEventsEventBusCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
