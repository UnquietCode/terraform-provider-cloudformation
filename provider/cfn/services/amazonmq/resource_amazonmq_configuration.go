// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package amazonmq

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAmazonMQConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAmazonMQConfigurationCreate,
		Read:   resourceAmazonMQConfigurationRead,
		Update: resourceAmazonMQConfigurationUpdate,
		Delete: resourceAmazonMQConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"engine_version": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"engine_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTagsEntry(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAmazonMQConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AmazonMQ::Configuration", data, meta)
}

func resourceAmazonMQConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AmazonMQ::Configuration", data, meta)
}

func resourceAmazonMQConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AmazonMQ::Configuration", data, meta)
}

func resourceAmazonMQConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AmazonMQ::Configuration", data, meta)
}