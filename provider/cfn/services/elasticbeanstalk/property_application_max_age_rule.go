// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationMaxAgeRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"delete_source_from_s3": {
				Type: schema.TypeBool,
				Required: false,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"max_age_in_days": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}