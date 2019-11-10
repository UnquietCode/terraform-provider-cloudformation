// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html

package sns

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSNSSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourceSNSSubscriptionCreate,
		Read:   resourceSNSSubscriptionRead,
		Update: resourceSNSSubscriptionUpdate,
		Delete: resourceSNSSubscriptionDelete,

		Schema: map[string]*schema.Schema{
			"delivery_policy": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"endpoint": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"filter_policy": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"protocol": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"raw_message_delivery": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"region": {
				Type: schema.TypeString,
				Optional: true,
			},
			"topic_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSNSSubscriptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SNS::Subscription", ResourceSNSSubscription(), data, meta)
}

func resourceSNSSubscriptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SNS::Subscription", ResourceSNSSubscription(), data, meta)
}

func resourceSNSSubscriptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SNS::Subscription", ResourceSNSSubscription(), data, meta)
}

func resourceSNSSubscriptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SNS::Subscription", data, meta)
}