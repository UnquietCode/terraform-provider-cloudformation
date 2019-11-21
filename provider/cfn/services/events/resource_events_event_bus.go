// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html

package events

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eventsEventBusType string = "AWS::Events::EventBus"

func ResourceEventsEventBus() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEventsEventBusRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eventsEventBusType, ResourceEventsEventBus(), data, meta)
}

func resourceEventsEventBusCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eventsEventBusType, ResourceEventsEventBus(), data, meta)
}

func resourceEventsEventBusUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eventsEventBusType, ResourceEventsEventBus(), data, meta)
}

func resourceEventsEventBusDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eventsEventBusType, data, meta)
}

func resourceEventsEventBusCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eventsEventBusType, data, meta)
}
