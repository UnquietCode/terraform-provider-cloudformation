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

func propertyRemediationParameterValue() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_value": {
				Type: schema.TypeList,
				Elem: propertyResourceValue(),
				Required: false,
				MaxItems: 1,
			},
			"static_value": {
				Type: schema.TypeList,
				Elem: propertyStaticValue(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}