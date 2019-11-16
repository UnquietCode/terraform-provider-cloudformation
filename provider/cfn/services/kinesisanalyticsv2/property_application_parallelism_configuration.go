// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-parallelismconfiguration.html

package kinesisanalyticsv2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationParallelismConfiguration(extras...string) *schema.Resource {
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
			"configuration_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"parallelism_per_kpu": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"auto_scaling_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"parallelism": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
