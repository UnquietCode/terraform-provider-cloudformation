// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package lambda

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCode() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"s3_bucket": {
				Type: schema.TypeString,
				Required: false,
			},
			"s3_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"s3_object_version": {
				Type: schema.TypeString,
				Required: false,
			},
			"zip_file": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}