// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package elasticbeanstalk

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationVersionLifecycleConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_age_rule": {
				Type: schema.TypeList,
				Elem: propertyMaxAgeRule(),
				Required: false,
				MaxItems: 1,
			},
			"max_count_rule": {
				Type: schema.TypeList,
				Elem: propertyMaxCountRule(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}