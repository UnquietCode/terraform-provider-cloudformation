// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolVerificationMessageTemplate() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"email_message_by_link": {
				Type: schema.TypeString,
				Required: false,
			},
			"email_message": {
				Type: schema.TypeString,
				Required: false,
			},
			"sms_message": {
				Type: schema.TypeString,
				Required: false,
			},
			"email_subject": {
				Type: schema.TypeString,
				Required: false,
			},
			"default_email_option": {
				Type: schema.TypeString,
				Required: false,
			},
			"email_subject_by_link": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}