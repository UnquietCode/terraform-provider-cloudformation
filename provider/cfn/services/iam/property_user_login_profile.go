// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package iam

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserLoginProfile() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"password": {
				Type: schema.TypeString,
				Required: true,
			},
			"password_reset_required": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}