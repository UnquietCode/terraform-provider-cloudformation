// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationType() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"actions": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentAccountTakeoverActionsType(),
				Required: true,
				MaxItems: 1,
			},
			"notify_configuration": {
				Type: schema.TypeList,
				Elem: propertyUserPoolRiskConfigurationAttachmentNotifyConfigurationType(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}