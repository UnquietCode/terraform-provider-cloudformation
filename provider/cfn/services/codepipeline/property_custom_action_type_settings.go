// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCustomActionTypeSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"entity_url_template": {
				Type: schema.TypeString,
				Required: false,
			},
			"execution_url_template": {
				Type: schema.TypeString,
				Required: false,
			},
			"revision_url_template": {
				Type: schema.TypeString,
				Required: false,
			},
			"third_party_configuration_url": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}