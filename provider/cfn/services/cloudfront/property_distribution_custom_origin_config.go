// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDistributionCustomOriginConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"origin_read_timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
			"https_port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"origin_keepalive_timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
			"origin_ssl_protocols": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"http_port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"origin_protocol_policy": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}