// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codecommit

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRepositoryRepositoryTrigger() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"events": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"branches": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"custom_data": {
				Type: schema.TypeString,
				Required: false,
			},
			"destination_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}