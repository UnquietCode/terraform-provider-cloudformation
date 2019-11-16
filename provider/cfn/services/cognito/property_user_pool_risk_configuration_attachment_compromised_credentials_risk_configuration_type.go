// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpoolriskconfigurationattachment-compromisedcredentialsriskconfigurationtype.html

package cognito

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var userPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationTypeProperties map[string]string = map[string]string{
	"actions": "Actions",
	"event_filter": "EventFilter",
}

func propertyUserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfigurationType(extras...string) *schema.Resource {
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
			"actions": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentCompromisedCredentialsActionsType(),
				Required: true,
				MaxItems: 1,
			},
			"event_filter": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}
