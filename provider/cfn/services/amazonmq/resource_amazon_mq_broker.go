// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package amazonmq

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAmazonMQBroker() *schema.Resource {
	return &schema.Resource{
		Create: resourceAmazonMQBrokerCreate,
		Read:   resourceAmazonMQBrokerRead,
		Update: resourceAmazonMQBrokerUpdate,
		Delete: resourceAmazonMQBrokerDelete,

		Schema: map[string]*schema.Schema{
			"security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"engine_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"configuration": {
				Type: schema.TypeList,
				Elem: propertyConfigurationId(),
				Required: false,
				MaxItems: 1,
			},
			"maintenance_window_start_time": {
				Type: schema.TypeList,
				Elem: propertyMaintenanceWindow(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"host_instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"auto_minor_version_upgrade": {
				Type: schema.TypeBool,
				Required: true,
			},
			"users": {
				Type: schema.TypeList,
				Elem: propertyUser(),
				Required: true,
			},
			"logs": {
				Type: schema.TypeList,
				Elem: propertyLogList(),
				Required: false,
				MaxItems: 1,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"broker_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deployment_mode": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"engine_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"publicly_accessible": {
				Type: schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"encryption_options": {
				Type: schema.TypeList,
				Elem: propertyEncryptionOptions(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyTagsEntry(),
				Required: false,
			},
		},
	}
}

func resourceAmazonMQBrokerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AmazonMQ::Broker", data, meta)
}

func resourceAmazonMQBrokerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AmazonMQ::Broker", data, meta)
}

func resourceAmazonMQBrokerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AmazonMQ::Broker", data, meta)
}

func resourceAmazonMQBrokerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AmazonMQ::Broker", data, meta)
}