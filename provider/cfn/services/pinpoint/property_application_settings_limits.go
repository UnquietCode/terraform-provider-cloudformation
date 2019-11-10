// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationSettingsLimits() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"daily": {
				Type: schema.TypeInt,
				Required: false,
			},
			"maximum_duration": {
				Type: schema.TypeInt,
				Required: false,
			},
			"total": {
				Type: schema.TypeInt,
				Required: false,
			},
			"messages_per_second": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}