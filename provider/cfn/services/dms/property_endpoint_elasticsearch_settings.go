// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEndpointElasticsearchSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"endpoint_uri": {
				Type: schema.TypeString,
				Required: false,
			},
			"full_load_error_percentage": {
				Type: schema.TypeInt,
				Required: false,
			},
			"error_retry_duration": {
				Type: schema.TypeInt,
				Required: false,
			},
			"service_access_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}