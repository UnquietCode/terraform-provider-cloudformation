// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLaunchTemplateLaunchTemplateData(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateTagSpecification(),
				Required: false,
			},
			"user_data": {
				Type: schema.TypeString,
				Required: false,
			},
			"block_device_mappings": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateBlockDeviceMapping(),
				Required: false,
			},
			"iam_instance_profile": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateIamInstanceProfile(),
				Required: false,
				MaxItems: 1,
			},
			"kernel_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Required: false,
			},
			"elastic_gpu_specifications": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateElasticGpuSpecification(),
				Required: false,
			},
			"elastic_inference_accelerators": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateLaunchTemplateElasticInferenceAccelerator(),
				Required: false,
			},
			"placement": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplatePlacement(),
				Required: false,
				MaxItems: 1,
			},
			"network_interfaces": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateNetworkInterface(),
				Required: false,
			},
			"image_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"monitoring": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateMonitoring(),
				Required: false,
				MaxItems: 1,
			},
			"hibernation_options": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateHibernationOptions(),
				Required: false,
				MaxItems: 1,
			},
			"license_specifications": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateLicenseSpecification(),
				Required: false,
			},
			"instance_initiated_shutdown_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"cpu_options": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateCpuOptions(),
				Required: false,
				MaxItems: 1,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"key_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"disable_api_termination": {
				Type: schema.TypeBool,
				Required: false,
			},
			"instance_market_options": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateInstanceMarketOptions(),
				Required: false,
				MaxItems: 1,
			},
			"ram_disk_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"capacity_reservation_specification": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateCapacityReservationSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"credit_specification": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateCreditSpecification(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}