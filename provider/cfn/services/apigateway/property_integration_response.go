// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyIntegrationResponse() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content_handling": {
				Type: schema.TypeString,
				Required: false,
			},
			"response_parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"response_templates": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"selection_pattern": {
				Type: schema.TypeString,
				Required: false,
			},
			"status_code": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}