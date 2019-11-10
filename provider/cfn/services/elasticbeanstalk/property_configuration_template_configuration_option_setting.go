// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyConfigurationTemplateConfigurationOptionSetting() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type: schema.TypeString,
				Required: true,
			},
			"option_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"resource_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"value": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}