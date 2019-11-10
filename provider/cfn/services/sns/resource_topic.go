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

func ResourceTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceTopicCreate,
		Read:   resourceTopicRead,
		Update: resourceTopicUpdate,
		Delete: resourceTopicDelete,

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"kms_master_key_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"subscription": {
				Type: schema.TypeList,
				Elem: propertyTopicSubscription(),
				Required: false,
			},
			"topic_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceTopicCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::SNS::Topic", data, meta)
}

func resourceTopicRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::SNS::Topic", data, meta)
}

func resourceTopicUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::SNS::Topic", data, meta)
}

func resourceTopicDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::SNS::Topic", data, meta)
}