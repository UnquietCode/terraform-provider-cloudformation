// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-verificationmessagetemplate.html

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
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"email_message_by_link": {
				Type: schema.TypeString,
				Optional: true,
			},
			"email_message": {
				Type: schema.TypeString,
				Optional: true,
			},
			"sms_message": {
				Type: schema.TypeString,
				Optional: true,
			},
			"email_subject": {
				Type: schema.TypeString,
				Optional: true,
			},
			"default_email_option": {
				Type: schema.TypeString,
				Optional: true,
			},
			"email_subject_by_link": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
