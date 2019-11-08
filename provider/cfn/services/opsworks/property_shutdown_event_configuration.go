// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyShutdownEventConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delay_until_elb_connections_drained": {
				Type: schema.TypeBool,
				Required: false,
			},
			"execution_timeout": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}