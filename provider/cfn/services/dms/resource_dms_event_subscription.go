// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

const dMSEventSubscriptionType string = "AWS::DMS::EventSubscription"

var dMSEventSubscriptionProperties map[string]string = map[string]string{
	"source_type": "SourceType",
	"event_categories": "EventCategories",
	"enabled": "Enabled",
	"subscription_name": "SubscriptionName",
	"sns_topic_arn": "SnsTopicArn",
	"source_ids": "SourceIds",
	"tags": "Tags",
}

func ResourceDMSEventSubscription() *schema.Resource {
	return &schema.Resource{
		Exists: resourceDMSEventSubscriptionExists,
		Read: resourceDMSEventSubscriptionRead,
		Create: resourceDMSEventSubscriptionCreate,
		Update: resourceDMSEventSubscriptionUpdate,
		Delete: resourceDMSEventSubscriptionDelete,
		CustomizeDiff: resourceDMSEventSubscriptionCustomizeDiff,
		
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
			},
			"sns_topic_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDMSEventSubscriptionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceDMSEventSubscriptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(dMSEventSubscriptionType, ResourceDMSEventSubscription(), data, meta)
}

func resourceDMSEventSubscriptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(dMSEventSubscriptionType, ResourceDMSEventSubscription(), data, dMSEventSubscriptionProperties, meta)
}

func resourceDMSEventSubscriptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(dMSEventSubscriptionType, ResourceDMSEventSubscription(), data, dMSEventSubscriptionProperties, meta)
}

func resourceDMSEventSubscriptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(dMSEventSubscriptionType, data, meta)
}

func resourceDMSEventSubscriptionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(dMSEventSubscriptionType, data, meta)
}
