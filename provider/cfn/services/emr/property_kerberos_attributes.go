// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyKerberosAttributes() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ad_domain_join_password": {
				Type: schema.TypeString,
				Required: false,
			},
			"ad_domain_join_user": {
				Type: schema.TypeString,
				Required: false,
			},
			"cross_realm_trust_principal_password": {
				Type: schema.TypeString,
				Required: false,
			},
			"kdc_admin_password": {
				Type: schema.TypeString,
				Required: true,
			},
			"realm": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}