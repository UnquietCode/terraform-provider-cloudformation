// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFunction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"function_arn": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"function_configuration": {
				Type: schema.TypeList,
				Elem: propertyFunctionConfiguration(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}