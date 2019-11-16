// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-fleetlaunchtemplateconfigrequest.html

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var eC2FleetFleetLaunchTemplateConfigRequestProperties map[string]string = map[string]string{
	"launch_template_specification": "LaunchTemplateSpecification",
	"overrides": "Overrides",
}

func propertyEC2FleetFleetLaunchTemplateConfigRequest(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"launch_template_specification": {
				Type: schema.TypeSet,
				Elem: propertyEC2FleetFleetLaunchTemplateSpecificationRequest(),
				Optional: true,
				MaxItems: 1,
			},
			"overrides": {
				Type: schema.TypeList,
				Elem: propertyEC2FleetFleetLaunchTemplateOverridesRequest(),
				Optional: true,
			},
		},
	}
}
