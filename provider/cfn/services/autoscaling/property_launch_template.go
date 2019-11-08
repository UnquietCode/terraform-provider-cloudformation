// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLaunchTemplate() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"launch_template_specification": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateSpecification(),
				Required: true,
				MaxItems: 1,
			},
			"overrides": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateOverrides(),
				Required: false,
			},
		},
	}
}