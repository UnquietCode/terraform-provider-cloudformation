// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package gamelift

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyIpPermission() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_port": {
				Type: schema.TypeInt,
				Required: true,
			},
			"ip_range": {
				Type: schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type: schema.TypeString,
				Required: true,
			},
			"to_port": {
				Type: schema.TypeInt,
				Required: true,
			},
		},
	}
}