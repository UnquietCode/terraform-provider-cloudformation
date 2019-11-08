// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package sagemaker

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyProductionVariant() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"model_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"variant_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"initial_instance_count": {
				Type: schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"accelerator_type": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"initial_variant_weight": {
				Type: schema.TypeFloat,
				Required: true,
				ForceNew: true,
			},
		},
	}
}