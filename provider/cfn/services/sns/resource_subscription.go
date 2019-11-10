// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package sns

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubscriptionCreate,
		Read:   resourceSubscriptionRead,
		Update: resourceSubscriptionUpdate,
		Delete: resourceSubscriptionDelete,

		Schema: map[string]*schema.Schema{
			"delivery_policy": {
				Type: schema.TypeMap,
				Required: false,
			},
			"endpoint": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"filter_policy": {
				Type: schema.TypeMap,
				Required: false,
			},
			"protocol": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"raw_message_delivery": {
				Type: schema.TypeBool,
				Required: false,
			},
			"region": {
				Type: schema.TypeString,
				Required: false,
			},
			"topic_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSubscriptionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SNS::Subscription", data, meta)
}

func resourceSubscriptionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SNS::Subscription", data, meta)
}

func resourceSubscriptionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SNS::Subscription", data, meta)
}

func resourceSubscriptionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SNS::Subscription", data, meta)
}