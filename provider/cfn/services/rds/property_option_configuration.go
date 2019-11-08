// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package rds

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyOptionConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"db_security_group_memberships": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"option_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"option_settings": {
				Type: schema.TypeSet,
				Elem: propertyOptionSetting(),
				Required: false,
			},
			"option_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"vpc_security_group_memberships": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
		},
	}
}