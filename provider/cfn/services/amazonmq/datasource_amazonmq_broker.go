// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html

package amazonmq

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const amazonMQBrokerType string = "AWS::AmazonMQ::Broker"

func DatasourceAmazonMQBroker() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAmazonMQBrokerRead,
		
		Schema: map[string]*schema.Schema{
			"security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"storage_type": {
				Type: schema.TypeString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAmazonMQBrokerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amazonMQBrokerType, DatasourceAmazonMQBroker(), data, meta)
}
