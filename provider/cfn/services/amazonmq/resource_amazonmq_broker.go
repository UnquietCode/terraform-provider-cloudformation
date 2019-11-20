// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html

package amazonmq

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const amazonMQBrokerType string = "AWS::AmazonMQ::Broker"

func ResourceAmazonMQBroker() *schema.Resource {
	return &schema.Resource{
		Read: resourceAmazonMQBrokerRead,
		Create: resourceAmazonMQBrokerCreate,
		Update: resourceAmazonMQBrokerUpdate,
		Delete: resourceAmazonMQBrokerDelete,
		CustomizeDiff: resourceAmazonMQBrokerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"engine_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"configuration": {
				Type: schema.TypeSet,
				Elem: propertyBrokerConfigurationId(),
				Optional: true,
				MaxItems: 1,
			},
			"maintenance_window_start_time": {
				Type: schema.TypeSet,
				Elem: propertyBrokerMaintenanceWindow(),
				Optional: true,
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
				Type: schema.TypeSet,
				Elem: propertyBrokerLogList(),
				Optional: true,
				MaxItems: 1,
			},
			"subnet_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"broker_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"deployment_mode": {
				Type: schema.TypeString,
				Required: true,
			},
			"engine_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"publicly_accessible": {
				Type: schema.TypeBool,
				Required: true,
			},
			"encryption_options": {
				Type: schema.TypeSet,
				Elem: propertyBrokerEncryptionOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyBrokerTagsEntry(),
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

func resourceAmazonMQBrokerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amazonMQBrokerType, ResourceAmazonMQBroker(), data, meta)
}

func resourceAmazonMQBrokerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(amazonMQBrokerType, ResourceAmazonMQBroker(), data, meta)
}

func resourceAmazonMQBrokerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(amazonMQBrokerType, ResourceAmazonMQBroker(), data, meta)
}

func resourceAmazonMQBrokerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(amazonMQBrokerType, data, meta)
}

func resourceAmazonMQBrokerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(amazonMQBrokerType, data, meta)
}
