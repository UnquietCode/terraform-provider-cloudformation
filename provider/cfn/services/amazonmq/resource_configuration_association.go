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

func ResourceConfigurationAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationAssociationCreate,
		Read:   resourceConfigurationAssociationRead,
		Update: resourceConfigurationAssociationUpdate,
		Delete: resourceConfigurationAssociationDelete,

		Schema: map[string]*schema.Schema{
			"broker": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"configuration": {
				Type: schema.TypeList,
				Elem: propertyConfigurationAssociationConfigurationId(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceConfigurationAssociationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AmazonMQ::ConfigurationAssociation", data, meta)
}

func resourceConfigurationAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AmazonMQ::ConfigurationAssociation", data, meta)
}

func resourceConfigurationAssociationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AmazonMQ::ConfigurationAssociation", data, meta)
}

func resourceConfigurationAssociationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AmazonMQ::ConfigurationAssociation", data, meta)
}