// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-configuration.html

package amazonmq

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const amazonMQConfigurationType string = "AWS::AmazonMQ::Configuration"

func DatasourceAmazonMQConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAmazonMQConfigurationRead,
		
		Schema: map[string]*schema.Schema{
			"engine_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"engine_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"data": {
				Type: schema.TypeString,
				Required: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: propertyConfigurationTagsEntry(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceAmazonMQConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(amazonMQConfigurationType, DatasourceAmazonMQConfiguration(), data, meta)
}
