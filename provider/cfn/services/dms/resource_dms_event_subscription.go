// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDMSEventSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourceDMSEventSubscriptionCreate,
		Read:   resourceDMSEventSubscriptionRead,
		Update: resourceDMSEventSubscriptionUpdate,
		Delete: resourceDMSEventSubscriptionDelete,

		Schema: map[string]*schema.Schema{
			"source_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"event_categories": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"subscription_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"sns_topic_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceDMSEventSubscriptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::DMS::EventSubscription", data, meta)
}

func resourceDMSEventSubscriptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::DMS::EventSubscription", data, meta)
}

func resourceDMSEventSubscriptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::DMS::EventSubscription", data, meta)
}

func resourceDMSEventSubscriptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::DMS::EventSubscription", data, meta)
}