// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package cloudfront

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyStreamingDistributionConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"logging": {
				Type: schema.TypeList,
				Elem: propertyLogging(),
				Required: false,
				MaxItems: 1,
			},
			"comment": {
				Type: schema.TypeString,
				Required: true,
			},
			"price_class": {
				Type: schema.TypeString,
				Required: false,
			},
			"s3_origin": {
				Type: schema.TypeList,
				Elem: propertyS3Origin(),
				Required: true,
				MaxItems: 1,
			},
			"enabled": {
				Type: schema.TypeBool,
				Required: true,
			},
			"aliases": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"trusted_signers": {
				Type: schema.TypeList,
				Elem: propertyTrustedSigners(),
				Required: true,
				MaxItems: 1,
			},
		},
	}
}