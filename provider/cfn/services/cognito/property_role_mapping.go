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

func propertyRoleMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"ambiguous_role_resolution": {
				Type: schema.TypeString,
				Required: false,
			},
			"rules_configuration": {
				Type: schema.TypeList,
				Elem: propertyRulesConfigurationType(),
				Required: false,
				MaxItems: 1,
			},
			"identity_provider": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}