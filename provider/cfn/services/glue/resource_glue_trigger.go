// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-trigger.html

package glue

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const glueTriggerType string = "AWS::Glue::Trigger"

func ResourceGlueTrigger() *schema.Resource {
	return &schema.Resource{
		Read: resourceGlueTriggerRead,
		Create: resourceGlueTriggerCreate,
		Update: resourceGlueTriggerUpdate,
		Delete: resourceGlueTriggerDelete,
		CustomizeDiff: resourceGlueTriggerCustomizeDiff,
		
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
			},
			"schedule": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"predicate": {
				Type: schema.TypeSet,
				Elem: propertyTriggerPredicate(),
				Optional: true,
				MaxItems: 1,
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

func resourceGlueTriggerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(glueTriggerType, ResourceGlueTrigger(), data, meta)
}

func resourceGlueTriggerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(glueTriggerType, ResourceGlueTrigger(), data, meta)
}

func resourceGlueTriggerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(glueTriggerType, ResourceGlueTrigger(), data, meta)
}

func resourceGlueTriggerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(glueTriggerType, data, meta)
}

func resourceGlueTriggerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(glueTriggerType, data, meta)
}
