// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyOrganizationAggregationSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_aws_regions": {
				Type: schema.TypeBool,
				Required: false,
			},
			"aws_regions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}