// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolDeviceConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device_only_remembered_on_user_prompt": {
				Type: schema.TypeBool,
				Required: false,
			},
			"challenge_required_on_new_device": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}