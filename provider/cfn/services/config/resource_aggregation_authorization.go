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

func ResourceAggregationAuthorization() *schema.Resource {
	return &schema.Resource{
		Create: resourceAggregationAuthorizationCreate,
		Read:   resourceAggregationAuthorizationRead,
		Update: resourceAggregationAuthorizationUpdate,
		Delete: resourceAggregationAuthorizationDelete,

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

func resourceAggregationAuthorizationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::Config::AggregationAuthorization", data, meta)
}

func resourceAggregationAuthorizationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::Config::AggregationAuthorization", data, meta)
}

func resourceAggregationAuthorizationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::Config::AggregationAuthorization", data, meta)
}

func resourceAggregationAuthorizationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::Config::AggregationAuthorization", data, meta)
}