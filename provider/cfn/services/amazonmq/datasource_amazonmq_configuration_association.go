// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configurationassociation.html

package amazonmq

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const amazonMQConfigurationAssociationType string = "AWS::AmazonMQ::ConfigurationAssociation"

func DatasourceAmazonMQConfigurationAssociation() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAmazonMQConfigurationAssociationRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAmazonMQConfigurationAssociationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amazonMQConfigurationAssociationType, DatasourceAmazonMQConfigurationAssociation(), data, meta)
}
