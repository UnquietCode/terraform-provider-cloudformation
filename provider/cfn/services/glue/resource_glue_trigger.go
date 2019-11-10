// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceGlueTrigger() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlueTriggerCreate,
		Read:   resourceGlueTriggerRead,
		Update: resourceGlueTriggerUpdate,
		Delete: resourceGlueTriggerDelete,

		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"start_on_creation": {
				Type: schema.TypeBool,
				Required: false,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"actions": {
				Type: schema.TypeList,
				Elem: propertyTriggerAction(),
				Required: true,
			},
			"workflow_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"schedule": {
				Type: schema.TypeString,
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"predicate": {
				Type: schema.TypeList,
				Elem: propertyTriggerPredicate(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceGlueTriggerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Trigger", data, meta)
}

func resourceGlueTriggerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Trigger", data, meta)
}

func resourceGlueTriggerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Trigger", data, meta)
}

func resourceGlueTriggerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Trigger", data, meta)
}