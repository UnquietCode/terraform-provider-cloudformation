// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package batch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyComputeResources() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"spot_iam_fleet_role": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"maxv_cpus": {
				Type: schema.TypeInt,
				Required: true,
			},
			"bid_percentage": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"subnets": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"allocation_strategy": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"minv_cpus": {
				Type: schema.TypeInt,
				Required: true,
			},
			"launch_template": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateSpecification(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"image_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"instance_role": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
			},
			"ec2_key_pair": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"placement_group": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
				ForceNew: true,
			},
			"desiredv_cpus": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}