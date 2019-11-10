// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyConfigRuleSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"owner": {
				Type: schema.TypeString,
				Required: true,
			},
			"source_details": {
				Type: schema.TypeSet,
				Elem: propertyConfigRuleSourceDetail(),
				Required: false,
			},
			"source_identifier": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}