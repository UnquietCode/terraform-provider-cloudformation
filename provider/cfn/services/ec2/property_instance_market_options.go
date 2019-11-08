// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInstanceMarketOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"spot_options": {
				Type: schema.TypeList,
				Elem: propertySpotOptions(),
				Required: false,
				MaxItems: 1,
			},
			"market_type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}