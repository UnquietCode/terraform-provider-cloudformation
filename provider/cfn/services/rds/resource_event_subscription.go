// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEventSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventSubscriptionCreate,
		Read:   resourceEventSubscriptionRead,
		Update: resourceEventSubscriptionUpdate,
		Delete: resourceEventSubscriptionDelete,

		Schema: map[string]*schema.Schema{
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"event_categories": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"sns_topic_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"source_type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceEventSubscriptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::EventSubscription", data, meta)
}

func resourceEventSubscriptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::EventSubscription", data, meta)
}

func resourceEventSubscriptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::EventSubscription", data, meta)
}

func resourceEventSubscriptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::EventSubscription", data, meta)
}