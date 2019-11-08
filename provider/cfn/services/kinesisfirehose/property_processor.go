// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package kinesisfirehose

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyProcessor() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"parameters": {
				Type: schema.TypeSet,
				Elem: propertyProcessorParameter(),
				Required: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}