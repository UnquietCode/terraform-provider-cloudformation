// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package amazonmq

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMaintenanceWindow() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"day_of_week": {
				Type: schema.TypeString,
				Required: true,
			},
			"time_of_day": {
				Type: schema.TypeString,
				Required: true,
			},
			"time_zone": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}