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

func propertyLoadBalancersConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"classic_load_balancers_config": {
				Type: schema.TypeList,
				Elem: propertyClassicLoadBalancersConfig(),
				Required: false,
				MaxItems: 1,
			},
			"target_groups_config": {
				Type: schema.TypeList,
				Elem: propertyTargetGroupsConfig(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}