// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html

package cognito

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var userPoolEmailConfigurationProperties map[string]string = map[string]string{
	"reply_to_email_address": "ReplyToEmailAddress",
	"email_sending_account": "EmailSendingAccount",
	"source_arn": "SourceArn",
}

func propertyUserPoolEmailConfiguration(extras...string) *schema.Resource {
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
			"reply_to_email_address": {
				Type: schema.TypeString,
				Optional: true,
			},
			"email_sending_account": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
