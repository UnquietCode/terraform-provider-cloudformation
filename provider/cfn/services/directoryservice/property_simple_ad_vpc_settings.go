// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package directoryservice

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySimpleADVpcSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"subnet_ids": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				Set: schema.HashString,
			},
			"vpc_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}