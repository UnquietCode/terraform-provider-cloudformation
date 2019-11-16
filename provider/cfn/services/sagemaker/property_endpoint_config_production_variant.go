// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-productionvariant.html

package sagemaker

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEndpointConfigProductionVariant(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"model_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"variant_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"initial_instance_count": {
				Type: schema.TypeInt,
				Required: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"accelerator_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"initial_variant_weight": {
				Type: schema.TypeFloat,
				Required: true,
			},
		},
	}
}
