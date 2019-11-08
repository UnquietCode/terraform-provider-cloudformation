// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package managedblockchain

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyNetworkConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"framework_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"voting_policy": {
				Type: schema.TypeList,
				Elem: propertyVotingPolicy(),
				Required: true,
				MaxItems: 1,
			},
			"framework": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"network_framework_configuration": {
				Type: schema.TypeList,
				Elem: propertyNetworkFrameworkConfiguration(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}