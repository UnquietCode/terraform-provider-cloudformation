// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigConfigurationAggregator() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigConfigurationAggregatorCreate,
		Read:   resourceConfigConfigurationAggregatorRead,
		Update: resourceConfigConfigurationAggregatorUpdate,
		Delete: resourceConfigConfigurationAggregatorDelete,

		Schema: map[string]*schema.Schema{
			"account_aggregation_sources": {
				Type: schema.TypeList,
				Elem: propertyConfigurationAggregatorAccountAggregationSource(),
				Optional: true,
			},
			"configuration_aggregator_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organization_aggregation_source": {
				Type: schema.TypeList,
				Elem: propertyConfigurationAggregatorOrganizationAggregationSource(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}

func resourceConfigConfigurationAggregatorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigurationAggregator", data, meta)
}

func resourceConfigConfigurationAggregatorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigurationAggregator", data, meta)
}

func resourceConfigConfigurationAggregatorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigurationAggregator", data, meta)
}

func resourceConfigConfigurationAggregatorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigurationAggregator", data, meta)
}