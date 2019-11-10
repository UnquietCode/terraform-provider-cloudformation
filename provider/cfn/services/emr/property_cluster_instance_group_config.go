// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyClusterInstanceGroupConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_scaling_policy": {
				Type: schema.TypeList,
				Elem: propertyClusterAutoScalingPolicy(),
				Required: false,
				MaxItems: 1,
			},
			"bid_price": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"configurations": {
				Type: schema.TypeSet,
				Elem: propertyClusterConfiguration(),
				Required: false,
				ForceNew: true,
			},
			"ebs_configuration": {
				Type: schema.TypeList,
				Elem: propertyClusterEbsConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"instance_count": {
				Type: schema.TypeInt,
				Required: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"market": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}