// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html

package amazonmq

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAmazonMQConfiguration() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAmazonMQConfigurationExists,
		Read: resourceAmazonMQConfigurationRead,
		Create: resourceAmazonMQConfigurationCreate,
		Update: resourceAmazonMQConfigurationUpdate,
		Delete: resourceAmazonMQConfigurationDelete,
		CustomizeDiff: resourceAmazonMQConfigurationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"engine_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"data": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTagsEntry(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAmazonMQConfigurationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAmazonMQConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AmazonMQ::Configuration", ResourceAmazonMQConfiguration(), data, meta)
}

func resourceAmazonMQConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AmazonMQ::Configuration", ResourceAmazonMQConfiguration(), data, meta)
}

func resourceAmazonMQConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AmazonMQ::Configuration", ResourceAmazonMQConfiguration(), data, meta)
}

func resourceAmazonMQConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AmazonMQ::Configuration", data, meta)
}

func resourceAmazonMQConfigurationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::AmazonMQ::Configuration", data, meta)
}

