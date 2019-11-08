// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceConfigAggregationAuthorization() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigAggregationAuthorizationCreate,
		Read:   resourceConfigAggregationAuthorizationRead,
		Update: resourceConfigAggregationAuthorizationUpdate,
		Delete: resourceConfigAggregationAuthorizationDelete,

		Schema: map[string]*schema.Schema{
			"authorized_account_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"authorized_aws_region": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConfigAggregationAuthorizationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::AggregationAuthorization", data, meta)
}

func resourceConfigAggregationAuthorizationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::AggregationAuthorization", data, meta)
}

func resourceConfigAggregationAuthorizationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::AggregationAuthorization", data, meta)
}

func resourceConfigAggregationAuthorizationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::AggregationAuthorization", data, meta)
}