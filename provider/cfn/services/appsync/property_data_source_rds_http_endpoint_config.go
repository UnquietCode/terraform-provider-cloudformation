// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDataSourceRdsHttpEndpointConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"aws_region": {
				Type: schema.TypeString,
				Required: true,
			},
			"schema": {
				Type: schema.TypeString,
				Required: false,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"db_cluster_identifier": {
				Type: schema.TypeString,
				Required: true,
			},
			"aws_secret_store_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}