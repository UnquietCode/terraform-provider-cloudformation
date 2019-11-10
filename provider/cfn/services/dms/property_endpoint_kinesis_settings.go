// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEndpointKinesisSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"message_format": {
				Type: schema.TypeString,
				Required: false,
			},
			"stream_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"service_access_role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}