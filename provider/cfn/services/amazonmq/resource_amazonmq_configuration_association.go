// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html

package amazonmq

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAmazonMQConfigurationAssociation() *schema.Resource {
	return &schema.Resource{
		Exists: resourceAmazonMQConfigurationAssociationExists,
		Read: resourceAmazonMQConfigurationAssociationRead,
		Create: resourceAmazonMQConfigurationAssociationCreate,
		Update: resourceAmazonMQConfigurationAssociationUpdate,
		Delete: resourceAmazonMQConfigurationAssociationDelete,
		CustomizeDiff: resourceAmazonMQConfigurationAssociationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"broker": {
				Type: schema.TypeString,
				Required: true,
			},
			"configuration": {
				Type: schema.TypeList,
				Elem: propertyConfigurationAssociationConfigurationId(),
				Required: true,
				MaxItems: 1,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAmazonMQConfigurationAssociationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceAmazonMQConfigurationAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AmazonMQ::ConfigurationAssociation", ResourceAmazonMQConfigurationAssociation(), data, meta)
}

func resourceAmazonMQConfigurationAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AmazonMQ::ConfigurationAssociation", ResourceAmazonMQConfigurationAssociation(), data, meta)
}

func resourceAmazonMQConfigurationAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AmazonMQ::ConfigurationAssociation", ResourceAmazonMQConfigurationAssociation(), data, meta)
}

func resourceAmazonMQConfigurationAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AmazonMQ::ConfigurationAssociation", data, meta)
}

func resourceAmazonMQConfigurationAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
