// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertySchemaAttribute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"developer_only_attribute": {
				Type: schema.TypeBool,
				Required: false,
			},
			"mutable": {
				Type: schema.TypeBool,
				Required: false,
			},
			"attribute_data_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"string_attribute_constraints": {
				Type: schema.TypeList,
				Elem: propertyStringAttributeConstraints(),
				Required: false,
				MaxItems: 1,
			},
			"required": {
				Type: schema.TypeBool,
				Required: false,
			},
			"number_attribute_constraints": {
				Type: schema.TypeList,
				Elem: propertyNumberAttributeConstraints(),
				Required: false,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}