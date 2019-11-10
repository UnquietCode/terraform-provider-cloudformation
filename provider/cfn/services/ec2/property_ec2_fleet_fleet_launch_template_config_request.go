// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEC2FleetFleetLaunchTemplateConfigRequest() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"launch_template_specification": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetFleetLaunchTemplateSpecificationRequest(),
				Required: false,
				MaxItems: 1,
			},
			"overrides": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetFleetLaunchTemplateOverridesRequest(),
				Required: false,
			},
		},
	}
}