// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package pinpoint

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCampaignEmailMessage() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"from_address": {
				Type: schema.TypeString,
				Required: false,
			},
			"html_body": {
				Type: schema.TypeString,
				Required: false,
			},
			"title": {
				Type: schema.TypeString,
				Required: false,
			},
			"body": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}