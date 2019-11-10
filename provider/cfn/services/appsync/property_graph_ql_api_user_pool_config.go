// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package appsync

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyGraphQLApiUserPoolConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"app_id_client_regex": {
				Type: schema.TypeString,
				Required: false,
			},
			"user_pool_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"aws_region": {
				Type: schema.TypeString,
				Required: false,
			},
			"default_action": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}