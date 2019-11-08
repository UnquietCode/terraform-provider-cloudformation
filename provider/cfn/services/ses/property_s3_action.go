// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ses

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyS3Action() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"kms_key_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"topic_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"object_key_prefix": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}