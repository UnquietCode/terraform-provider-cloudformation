// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMemberMemberFabricConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"admin_username": {
				Type: schema.TypeString,
				Required: true,
			},
			"admin_password": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}