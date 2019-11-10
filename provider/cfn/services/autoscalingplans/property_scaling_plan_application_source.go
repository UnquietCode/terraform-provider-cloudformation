// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscalingplans

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyScalingPlanApplicationSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_formation_stack_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"tag_filters": {
				Type: schema.TypeList,
				Elem: propertyScalingPlanTagFilter(),
				Required: false,
			},
		},
	}
}