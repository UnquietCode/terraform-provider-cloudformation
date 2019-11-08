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

func propertyNotifyConfigurationType() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"block_email": {
				Type: schema.TypeList,
				Elem: propertyNotifyEmailType(),
				Required: false,
				MaxItems: 1,
			},
			"reply_to": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"no_action_email": {
				Type: schema.TypeList,
				Elem: propertyNotifyEmailType(),
				Required: false,
				MaxItems: 1,
			},
			"from": {
				Type: schema.TypeString,
				Required: false,
			},
			"mfa_email": {
				Type: schema.TypeList,
				Elem: propertyNotifyEmailType(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}