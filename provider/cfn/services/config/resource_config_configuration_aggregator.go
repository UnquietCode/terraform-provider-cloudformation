// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceConfigConfigurationAggregatorExists,
		Read:   resourceConfigConfigurationAggregatorRead,
		Create: resourceConfigConfigurationAggregatorCreate,
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
			},
			"organization_aggregation_source": {
				Type: schema.TypeList,
				Elem: propertyConfigurationAggregatorOrganizationAggregationSource(),
				Optional: true,
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

func resourceConfigConfigurationAggregatorExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceConfigConfigurationAggregatorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigurationAggregator", ResourceConfigConfigurationAggregator(), data, meta)
}

func resourceConfigConfigurationAggregatorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigurationAggregator", ResourceConfigConfigurationAggregator(), data, meta)
}

func resourceConfigConfigurationAggregatorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigurationAggregator", ResourceConfigConfigurationAggregator(), data, meta)
}

func resourceConfigConfigurationAggregatorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigurationAggregator", data, meta)
}