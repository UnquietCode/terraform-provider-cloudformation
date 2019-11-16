// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const configAggregationAuthorizationType string = "AWS::Config::AggregationAuthorization"

var configAggregationAuthorizationProperties map[string]string = map[string]string{
	"authorized_account_id": "AuthorizedAccountId",
	"authorized_aws_region": "AuthorizedAwsRegion",
}

func ResourceConfigAggregationAuthorization() *schema.Resource {
	return &schema.Resource{
		Exists: resourceConfigAggregationAuthorizationExists,
		Read: resourceConfigAggregationAuthorizationRead,
		Create: resourceConfigAggregationAuthorizationCreate,
		Update: resourceConfigAggregationAuthorizationUpdate,
		Delete: resourceConfigAggregationAuthorizationDelete,
		CustomizeDiff: resourceConfigAggregationAuthorizationCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"authorized_account_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorized_aws_region": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConfigAggregationAuthorizationExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceConfigAggregationAuthorizationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(configAggregationAuthorizationType, ResourceConfigAggregationAuthorization(), data, meta)
}

func resourceConfigAggregationAuthorizationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(configAggregationAuthorizationType, ResourceConfigAggregationAuthorization(), data, configAggregationAuthorizationProperties, meta)
}

func resourceConfigAggregationAuthorizationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(configAggregationAuthorizationType, ResourceConfigAggregationAuthorization(), data, configAggregationAuthorizationProperties, meta)
}

func resourceConfigAggregationAuthorizationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(configAggregationAuthorizationType, data, meta)
}

func resourceConfigAggregationAuthorizationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(configAggregationAuthorizationType, data, meta)
}
