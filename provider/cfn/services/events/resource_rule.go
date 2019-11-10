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

func ResourceRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceRuleCreate,
		Read:   resourceRuleRead,
		Update: resourceRuleUpdate,
		Delete: resourceRuleDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"event_pattern": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"schedule_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"state": {
				Type: schema.TypeString,
				Required: false,
			},
			"targets": {
				Type: schema.TypeSet,
				Elem: propertyRuleTarget(),
				Required: false,
			},
		},
	}
}

func resourceRuleCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Events::Rule", data, meta)
}

func resourceRuleRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Events::Rule", data, meta)
}

func resourceRuleUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Events::Rule", data, meta)
}

func resourceRuleDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Events::Rule", data, meta)
}