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

func propertyRdsDbInstance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"db_password": {
				Type: schema.TypeString,
				Required: true,
			},
			"db_user": {
				Type: schema.TypeString,
				Required: true,
			},
			"rds_db_instance_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}