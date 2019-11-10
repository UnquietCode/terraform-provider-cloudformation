// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package fsx

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFileSystemSelfManagedActiveDirectoryConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"file_system_administrators_group": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"user_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"domain_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"organizational_unit_distinguished_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"dns_ips": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"password": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}