// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigurationAggregator() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigurationAggregatorCreate,
		Read:   resourceConfigurationAggregatorRead,
		Update: resourceConfigurationAggregatorUpdate,
		Delete: resourceConfigurationAggregatorDelete,

		Schema: map[string]*schema.Schema{
			"account_aggregation_sources": {
				Type: schema.TypeList,
				Elem: propertyConfigurationAggregatorAccountAggregationSource(),
				Required: false,
			},
			"configuration_aggregator_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organization_aggregation_source": {
				Type: schema.TypeList,
				Elem: propertyConfigurationAggregatorOrganizationAggregationSource(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}

func resourceConfigurationAggregatorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::ConfigurationAggregator", data, meta)
}

func resourceConfigurationAggregatorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::ConfigurationAggregator", data, meta)
}

func resourceConfigurationAggregatorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::ConfigurationAggregator", data, meta)
}

func resourceConfigurationAggregatorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::ConfigurationAggregator", data, meta)
}