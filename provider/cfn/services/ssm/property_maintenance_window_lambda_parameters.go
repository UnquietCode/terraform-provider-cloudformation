// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyMaintenanceWindowLambdaParameters() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_context": {
				Type: schema.TypeString,
				Required: false,
			},
			"qualifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"payload": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}