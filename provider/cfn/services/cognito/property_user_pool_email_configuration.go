// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package cognito

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyUserPoolEmailConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"reply_to_email_address": {
				Type: schema.TypeString,
				Required: false,
			},
			"email_sending_account": {
				Type: schema.TypeString,
				Required: false,
			},
			"source_arn": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}