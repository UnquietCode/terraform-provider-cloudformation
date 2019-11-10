// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html

package sns

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceSNSTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceSNSTopicCreate,
		Read:   resourceSNSTopicRead,
		Update: resourceSNSTopicUpdate,
		Delete: resourceSNSTopicDelete,

		Schema: map[string]*schema.Schema{
			"topic_name": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
		},
	}
}

func resourceSNSTopicCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SNS::Topic", data, meta)
}

func resourceSNSTopicRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SNS::Topic", data, meta)
}

func resourceSNSTopicUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SNS::Topic", data, meta)
}

func resourceSNSTopicDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SNS::Topic", data, meta)
}