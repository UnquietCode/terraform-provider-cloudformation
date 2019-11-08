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

func propertyMaxCountRule() *schema.Resource {
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
			"max_count": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}