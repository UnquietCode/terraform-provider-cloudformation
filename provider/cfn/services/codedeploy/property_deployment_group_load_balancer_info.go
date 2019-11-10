// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeploymentGroupLoadBalancerInfo() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"elb_info_list": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentGroupELBInfo(),
				Required: false,
			},
			"target_group_info_list": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentGroupTargetGroupInfo(),
				Required: false,
			},
		},
	}
}