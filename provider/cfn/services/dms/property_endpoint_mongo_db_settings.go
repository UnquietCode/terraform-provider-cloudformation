// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package dms

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEndpointMongoDbSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auth_source": {
				Type: schema.TypeString,
				Required: false,
			},
			"auth_mechanism": {
				Type: schema.TypeString,
				Required: false,
			},
			"username": {
				Type: schema.TypeString,
				Required: false,
			},
			"docs_to_investigate": {
				Type: schema.TypeString,
				Required: false,
			},
			"server_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"port": {
				Type: schema.TypeInt,
				Required: false,
			},
			"extract_doc_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"database_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"auth_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"password": {
				Type: schema.TypeString,
				Required: false,
			},
			"nesting_level": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}