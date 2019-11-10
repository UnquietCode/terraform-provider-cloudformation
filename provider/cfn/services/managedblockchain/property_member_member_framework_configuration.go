// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMemberMemberFrameworkConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"member_fabric_configuration": {
				Type: schema.TypeList,
				Elem: propertyMemberMemberFabricConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}