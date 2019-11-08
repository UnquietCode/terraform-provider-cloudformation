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

func propertyConnectionInput() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"connection_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"match_criteria": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"physical_connection_requirements": {
				Type: schema.TypeList,
				Elem: propertyPhysicalConnectionRequirements(),
				Required: false,
				MaxItems: 1,
			},
			"connection_properties": {
				Type: schema.TypeMap,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}