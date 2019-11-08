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

func propertyAccountTakeoverActionsType() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"high_action": {
				Type: schema.TypeList,
				Elem: propertyAccountTakeoverActionType(),
				Required: false,
				MaxItems: 1,
			},
			"low_action": {
				Type: schema.TypeList,
				Elem: propertyAccountTakeoverActionType(),
				Required: false,
				MaxItems: 1,
			},
			"medium_action": {
				Type: schema.TypeList,
				Elem: propertyAccountTakeoverActionType(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}