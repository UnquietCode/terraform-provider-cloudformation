// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySsmControls() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"error_percentage": {
				Type: schema.TypeInt,
				Required: false,
			},
			"concurrent_execution_rate_percentage": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}