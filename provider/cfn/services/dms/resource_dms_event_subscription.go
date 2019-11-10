// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-eventsubscription.html

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
				Optional: true,
			},
			"event_categories": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"subscription_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"sns_topic_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
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