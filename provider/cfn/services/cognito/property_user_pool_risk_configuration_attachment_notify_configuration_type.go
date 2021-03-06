// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-notifyconfigurationtype.html

package cognito

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolRiskConfigurationAttachmentNotifyConfigurationType(extras...string) *schema.Resource {
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
			"block_email": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentNotifyEmailType(),
				Optional: true,
				MaxItems: 1,
			},
			"reply_to": {
				Type: schema.TypeString,
				Optional: true,
			},
			"source_arn": {
				Type: schema.TypeString,
				Required: true,
			},
			"no_action_email": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentNotifyEmailType(),
				Optional: true,
				MaxItems: 1,
			},
			"from": {
				Type: schema.TypeString,
				Optional: true,
			},
			"mfa_email": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentNotifyEmailType(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
