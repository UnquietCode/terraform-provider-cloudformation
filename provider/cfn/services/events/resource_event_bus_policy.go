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

func ResourceEventBusPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventBusPolicyCreate,
		Read:   resourceEventBusPolicyRead,
		Update: resourceEventBusPolicyUpdate,
		Delete: resourceEventBusPolicyDelete,

		Schema: map[string]*schema.Schema{
			"event_bus_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"condition": {
				Type: schema.TypeList,
				Elem: propertyEventBusPolicyCondition(),
				Required: false,
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
		},
	}
}

func resourceEventBusPolicyCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Events::EventBusPolicy", data, meta)
}

func resourceEventBusPolicyRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Events::EventBusPolicy", data, meta)
}

func resourceEventBusPolicyUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Events::EventBusPolicy", data, meta)
}

func resourceEventBusPolicyDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Events::EventBusPolicy", data, meta)
}