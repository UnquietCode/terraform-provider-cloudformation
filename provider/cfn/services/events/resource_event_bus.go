// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package events

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEventBus() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventBusCreate,
		Read:   resourceEventBusRead,
		Update: resourceEventBusUpdate,
		Delete: resourceEventBusDelete,

		Schema: map[string]*schema.Schema{
			"event_source_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceEventBusCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Events::EventBus", data, meta)
}

func resourceEventBusRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Events::EventBus", data, meta)
}

func resourceEventBusUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Events::EventBus", data, meta)
}

func resourceEventBusDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Events::EventBus", data, meta)
}