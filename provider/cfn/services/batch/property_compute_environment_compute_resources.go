// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-computeenvironment-computeresources.html

package batch

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var computeEnvironmentComputeResourcesProperties map[string]string = map[string]string{
	"spot_iam_fleet_role": "SpotIamFleetRole",
	"maxv_cpus": "MaxvCpus",
	"bid_percentage": "BidPercentage",
	"security_group_ids": "SecurityGroupIds",
	"subnets": "Subnets",
	"type": "Type",
	"allocation_strategy": "AllocationStrategy",
	"minv_cpus": "MinvCpus",
	"launch_template": "LaunchTemplate",
	"image_id": "ImageId",
	"instance_role": "InstanceRole",
	"instance_types": "InstanceTypes",
	"ec2_key_pair": "Ec2KeyPair",
	"placement_group": "PlacementGroup",
	"tags": "Tags",
	"desiredv_cpus": "DesiredvCpus",
}

func propertyComputeEnvironmentComputeResources(extras...string) *schema.Resource {
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
			"spot_iam_fleet_role": {
				Type: schema.TypeString,
				Optional: true,
			},
			"maxv_cpus": {
				Type: schema.TypeInt,
				Required: true,
			},
			"bid_percentage": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"subnets": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"allocation_strategy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"minv_cpus": {
				Type: schema.TypeInt,
				Required: true,
			},
			"launch_template": {
				Type: schema.TypeSet,
				Elem: propertyComputeEnvironmentLaunchTemplateSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"image_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_role": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_types": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"ec2_key_pair": {
				Type: schema.TypeString,
				Optional: true,
			},
			"placement_group": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"desiredv_cpus": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
