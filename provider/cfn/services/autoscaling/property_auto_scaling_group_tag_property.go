// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyAutoScalingGroupTagProperty() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type: schema.TypeString,
				Required: true,
			},
			"propagate_at_launch": {
				Type: schema.TypeBool,
				Required: true,
			},
			"value": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}