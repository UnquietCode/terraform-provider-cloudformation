// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package greengrass

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyFunctionDefinitionEnvironment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"variables": {
				Type: schema.TypeMap,
				Required: false,
				ForceNew: true,
			},
			"execution": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionExecution(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"resource_access_policies": {
				Type: schema.TypeList,
				Elem: propertyFunctionDefinitionResourceAccessPolicy(),
				Required: false,
				ForceNew: true,
			},
			"access_sysfs": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
		},
	}
}