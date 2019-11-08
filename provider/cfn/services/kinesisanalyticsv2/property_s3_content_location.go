// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyS3ContentLocation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bucket_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"file_key": {
				Type: schema.TypeString,
				Required: false,
			},
			"object_version": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}