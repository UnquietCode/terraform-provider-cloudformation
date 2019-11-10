// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClientVpnEndpointConnectionLogOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloudwatch_log_stream": {
				Type: schema.TypeString,
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"cloudwatch_log_group": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}