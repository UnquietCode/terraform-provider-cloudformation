// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticloadbalancing

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLoadBalancerAccessLoggingPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"emit_interval": {
				Type: schema.TypeInt,
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"s3_bucket_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"s3_bucket_prefix": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}