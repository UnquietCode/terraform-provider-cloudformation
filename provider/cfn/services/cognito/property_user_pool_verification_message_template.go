// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolVerificationMessageTemplate(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
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