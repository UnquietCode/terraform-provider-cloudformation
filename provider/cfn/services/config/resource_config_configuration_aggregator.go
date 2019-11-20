// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const configConfigurationAggregatorType string = "AWS::Config::ConfigurationAggregator"

func ResourceConfigConfigurationAggregator() *schema.Resource {
	return &schema.Resource{
		Read: resourceConfigConfigurationAggregatorRead,
		Create: resourceConfigConfigurationAggregatorCreate,
		Update: resourceConfigConfigurationAggregatorUpdate,
		Delete: resourceConfigConfigurationAggregatorDelete,
		CustomizeDiff: resourceConfigConfigurationAggregatorCustomizeDiff,
		
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
				Type: schema.TypeSet,
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

func resourceConfigConfigurationAggregatorRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(configConfigurationAggregatorType, ResourceConfigConfigurationAggregator(), data, meta)
}

func resourceConfigConfigurationAggregatorCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(configConfigurationAggregatorType, ResourceConfigConfigurationAggregator(), data, meta)
}

func resourceConfigConfigurationAggregatorUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(configConfigurationAggregatorType, ResourceConfigConfigurationAggregator(), data, meta)
}

func resourceConfigConfigurationAggregatorDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(configConfigurationAggregatorType, data, meta)
}

func resourceConfigConfigurationAggregatorCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(configConfigurationAggregatorType, data, meta)
}
