// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cloudwatch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"metric_time_zone": {
				Type: schema.TypeString,
				Required: false,
			},
			"excluded_time_ranges": {
				Type: schema.TypeList,
				Elem: propertyRange(),
				Required: false,
			},
		},
	}
}