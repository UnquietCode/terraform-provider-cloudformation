// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
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

func DatasourceConfigAggregationAuthorization() *schema.Resource {
	return &schema.Resource{
		Read: datasourceConfigAggregationAuthorizationRead,
		
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
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceConfigAggregationAuthorizationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(configAggregationAuthorizationType, DatasourceConfigAggregationAuthorization(), data, meta)
}
