// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
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

var amazonMQBrokerProperties map[string]string = map[string]string{
	"security_groups": "SecurityGroups",
	"engine_version": "EngineVersion",
	"configuration": "Configuration",
	"maintenance_window_start_time": "MaintenanceWindowStartTime",
	"host_instance_type": "HostInstanceType",
	"auto_minor_version_upgrade": "AutoMinorVersionUpgrade",
	"users": "Users",
	"logs": "Logs",
	"subnet_ids": "SubnetIds",
	"broker_name": "BrokerName",
	"deployment_mode": "DeploymentMode",
	"engine_type": "EngineType",
	"publicly_accessible": "PubliclyAccessible",
	"encryption_options": "EncryptionOptions",
	"tags": "Tags",
}

func ResourceAmazonMQBroker() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAmazonMQBrokerExists,
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
				Type: schema.TypeList,
				Elem: propertyBrokerConfigurationId(),
				Optional: true,
				MaxItems: 1,
			},
			"maintenance_window_start_time": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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

func resourceAmazonMQBrokerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAmazonMQBrokerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amazonMQBrokerType, ResourceAmazonMQBroker(), data, meta)
}

func resourceAmazonMQBrokerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(amazonMQBrokerType, ResourceAmazonMQBroker(), data, amazonMQBrokerProperties, meta)
}

func resourceAmazonMQBrokerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(amazonMQBrokerType, ResourceAmazonMQBroker(), data, amazonMQBrokerProperties, meta)
}

func resourceAmazonMQBrokerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(amazonMQBrokerType, data, meta)
}

func resourceAmazonMQBrokerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(amazonMQBrokerType, data, meta)
}
