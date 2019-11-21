// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-aggregationauthorization.html

package config

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const configAggregationAuthorizationType string = "AWS::Config::AggregationAuthorization"

func ResourceConfigAggregationAuthorization() *schema.Resource {
	return &schema.Resource{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceConfigAggregationAuthorizationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(configAggregationAuthorizationType, ResourceConfigAggregationAuthorization(), data, meta)
}

func resourceConfigAggregationAuthorizationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(configAggregationAuthorizationType, ResourceConfigAggregationAuthorization(), data, meta)
}

func resourceConfigAggregationAuthorizationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(configAggregationAuthorizationType, ResourceConfigAggregationAuthorization(), data, meta)
}

func resourceConfigAggregationAuthorizationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(configAggregationAuthorizationType, data, meta)
}

func resourceConfigAggregationAuthorizationCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(configAggregationAuthorizationType, data, meta)
}
