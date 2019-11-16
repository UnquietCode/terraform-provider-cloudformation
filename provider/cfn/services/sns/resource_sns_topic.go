// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html

package sns

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const sNSTopicType string = "AWS::SNS::Topic"

var sNSTopicProperties map[string]string = map[string]string{
	"display_name": "DisplayName",
	"kms_master_key_id": "KmsMasterKeyId",
	"subscription": "Subscription",
	"topic_name": "TopicName",
}

func ResourceSNSTopic() *schema.Resource {
	return &schema.Resource{
		Exists: resourceSNSTopicExists,
		Read: resourceSNSTopicRead,
		Create: resourceSNSTopicCreate,
		Update: resourceSNSTopicUpdate,
		Delete: resourceSNSTopicDelete,
		CustomizeDiff: resourceSNSTopicCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"kms_master_key_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subscription": {
				Type: schema.TypeList,
				Elem: propertyTopicSubscription(),
				Optional: true,
			},
			"topic_name": {
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

func resourceSNSTopicExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceSNSTopicRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(sNSTopicType, ResourceSNSTopic(), data, meta)
}

func resourceSNSTopicCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(sNSTopicType, ResourceSNSTopic(), data, sNSTopicProperties, meta)
}

func resourceSNSTopicUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(sNSTopicType, ResourceSNSTopic(), data, sNSTopicProperties, meta)
}

func resourceSNSTopicDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(sNSTopicType, data, meta)
}

func resourceSNSTopicCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(sNSTopicType, data, meta)
}
