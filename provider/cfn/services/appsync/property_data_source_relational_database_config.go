// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDataSourceRelationalDatabaseConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"rds_http_endpoint_config": {
				Type: schema.TypeList,
				Elem: propertyDataSourceRdsHttpEndpointConfig(),
				Required: false,
				MaxItems: 1,
			},
			"relational_database_source_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}