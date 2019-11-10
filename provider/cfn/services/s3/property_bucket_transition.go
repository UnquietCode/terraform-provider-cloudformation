// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package s3

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBucketTransition() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"storage_class": {
				Type: schema.TypeString,
				Required: true,
			},
			"transition_date": {
				Type: schema.TypeString,
				Required: false,
			},
			"transition_in_days": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}