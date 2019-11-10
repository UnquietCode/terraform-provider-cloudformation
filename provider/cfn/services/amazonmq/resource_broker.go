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

func ResourceBroker() *schema.Resource {
	return &schema.Resource{
		Create: resourceBrokerCreate,
		Read:   resourceBrokerRead,
		Update: resourceBrokerUpdate,
		Delete: resourceBrokerDelete,

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
				Elem: propertyBrokerConfigurationId(),
				Required: false,
				MaxItems: 1,
			},
			"maintenance_window_start_time": {
				Type: schema.TypeList,
				Elem: propertyBrokerMaintenanceWindow(),
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
				Elem: propertyBrokerUser(),
				Required: true,
			},
			"logs": {
				Type: schema.TypeList,
				Elem: propertyBrokerLogList(),
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
				Elem: propertyBrokerEncryptionOptions(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyBrokerTagsEntry(),
				Required: false,
			},
		},
	}
}

func resourceBrokerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AmazonMQ::Broker", data, meta)
}

func resourceBrokerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AmazonMQ::Broker", data, meta)
}

func resourceBrokerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AmazonMQ::Broker", data, meta)
}

func resourceBrokerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AmazonMQ::Broker", data, meta)
}