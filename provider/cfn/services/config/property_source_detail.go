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

func propertySourceDetail() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"event_source": {
				Type: schema.TypeString,
				Required: true,
			},
			"maximum_execution_frequency": {
				Type: schema.TypeString,
				Required: false,
			},
			"message_type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}