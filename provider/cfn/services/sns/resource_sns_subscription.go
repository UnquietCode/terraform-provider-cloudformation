// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html

package sns

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sNSSubscriptionType string = "AWS::SNS::Subscription"

func ResourceSNSSubscription() *schema.Resource {
	return &schema.Resource{
		Read: resourceSNSSubscriptionRead,
		Create: resourceSNSSubscriptionCreate,
		Update: resourceSNSSubscriptionUpdate,
		Delete: resourceSNSSubscriptionDelete,
		CustomizeDiff: resourceSNSSubscriptionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"delivery_policy": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"endpoint": {
				Type: schema.TypeString,
				Optional: true,
			},
			"filter_policy": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"protocol": {
				Type: schema.TypeString,
				Required: true,
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

func resourceSNSSubscriptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sNSSubscriptionType, ResourceSNSSubscription(), data, meta)
}

func resourceSNSSubscriptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sNSSubscriptionType, ResourceSNSSubscription(), data, meta)
}

func resourceSNSSubscriptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sNSSubscriptionType, ResourceSNSSubscription(), data, meta)
}

func resourceSNSSubscriptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sNSSubscriptionType, data, meta)
}

func resourceSNSSubscriptionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sNSSubscriptionType, data, meta)
}
