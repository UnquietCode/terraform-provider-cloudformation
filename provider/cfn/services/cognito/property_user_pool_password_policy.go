// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"require_numbers": {
				Type: schema.TypeBool,
				Required: false,
			},
			"minimum_length": {
				Type: schema.TypeInt,
				Required: false,
			},
			"temporary_password_validity_days": {
				Type: schema.TypeInt,
				Required: false,
			},
			"require_uppercase": {
				Type: schema.TypeBool,
				Required: false,
			},
			"require_lowercase": {
				Type: schema.TypeBool,
				Required: false,
			},
			"require_symbols": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}