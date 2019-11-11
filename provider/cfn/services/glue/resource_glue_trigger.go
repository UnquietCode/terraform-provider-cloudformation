// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html

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
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"actions": {
				Type: schema.TypeList,
				Elem: propertyTriggerAction(),
				Required: true,
			},
			"workflow_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"schedule": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"predicate": {
				Type: schema.TypeList,
				Elem: propertyTriggerPredicate(),
				Optional: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceGlueTriggerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Glue::Trigger", ResourceGlueTrigger(), data, meta)
}

func resourceGlueTriggerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Glue::Trigger", ResourceGlueTrigger(), data, meta)
}

func resourceGlueTriggerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Glue::Trigger", ResourceGlueTrigger(), data, meta)
}

func resourceGlueTriggerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Glue::Trigger", data, meta)
}