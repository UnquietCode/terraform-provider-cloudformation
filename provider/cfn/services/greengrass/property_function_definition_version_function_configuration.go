// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFunctionDefinitionVersionFunctionConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"memory_size": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"pinned": {
				Type: schema.TypeBool,
				Required: false,
			},
			"exec_args": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"timeout": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"encoding_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"environment": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionVersionEnvironment(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"executable": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}