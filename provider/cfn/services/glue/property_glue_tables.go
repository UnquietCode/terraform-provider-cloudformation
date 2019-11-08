// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package glue

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyGlueTables() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"connection_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"table_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"catalog_id": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}