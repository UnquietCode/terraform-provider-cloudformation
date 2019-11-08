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

func propertyAdminCreateUserConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"invite_message_template": {
				Type: schema.TypeList,
				Elem: propertyInviteMessageTemplate(),
				Required: false,
				MaxItems: 1,
			},
			"unused_account_validity_days": {
				Type: schema.TypeInt,
				Required: false,
			},
			"allow_admin_create_user_only": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}