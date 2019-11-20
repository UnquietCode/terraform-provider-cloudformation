// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html

package amazonmq

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const amazonMQConfigurationAssociationType string = "AWS::AmazonMQ::ConfigurationAssociation"

func ResourceAmazonMQConfigurationAssociation() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeSet,
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

func resourceAmazonMQConfigurationAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amazonMQConfigurationAssociationType, ResourceAmazonMQConfigurationAssociation(), data, meta)
}

func resourceAmazonMQConfigurationAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(amazonMQConfigurationAssociationType, ResourceAmazonMQConfigurationAssociation(), data, meta)
}

func resourceAmazonMQConfigurationAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(amazonMQConfigurationAssociationType, ResourceAmazonMQConfigurationAssociation(), data, meta)
}

func resourceAmazonMQConfigurationAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(amazonMQConfigurationAssociationType, data, meta)
}

func resourceAmazonMQConfigurationAssociationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(amazonMQConfigurationAssociationType, data, meta)
}
