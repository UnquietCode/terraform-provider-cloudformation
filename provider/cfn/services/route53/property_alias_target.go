// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package route53

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAliasTarget() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dns_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"evaluate_target_health": {
				Type: schema.TypeBool,
				Required: false,
			},
			"hosted_zone_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}