// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRDSEventSubscription() *schema.Resource {
	return &schema.Resource{
		Exists: resourceRDSEventSubscriptionExists,
		Read: resourceRDSEventSubscriptionRead,
		Create: resourceRDSEventSubscriptionCreate,
		Update: resourceRDSEventSubscriptionUpdate,
		Delete: resourceRDSEventSubscriptionDelete,
		CustomizeDiff: resourceRDSEventSubscriptionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"event_categories": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
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
			"source_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRDSEventSubscriptionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceRDSEventSubscriptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::RDS::EventSubscription", ResourceRDSEventSubscription(), data, meta)
}

func resourceRDSEventSubscriptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::RDS::EventSubscription", ResourceRDSEventSubscription(), data, meta)
}

func resourceRDSEventSubscriptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::RDS::EventSubscription", ResourceRDSEventSubscription(), data, meta)
}

func resourceRDSEventSubscriptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::RDS::EventSubscription", data, meta)
}

func resourceRDSEventSubscriptionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
