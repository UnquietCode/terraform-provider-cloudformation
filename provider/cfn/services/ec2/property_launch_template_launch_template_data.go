// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var launchTemplateLaunchTemplateDataProperties map[string]string = map[string]string{
	"security_groups": "SecurityGroups",
	"tag_specifications": "TagSpecifications",
	"user_data": "UserData",
	"block_device_mappings": "BlockDeviceMappings",
	"iam_instance_profile": "IamInstanceProfile",
	"kernel_id": "KernelId",
	"ebs_optimized": "EbsOptimized",
	"elastic_gpu_specifications": "ElasticGpuSpecifications",
	"elastic_inference_accelerators": "ElasticInferenceAccelerators",
	"placement": "Placement",
	"network_interfaces": "NetworkInterfaces",
	"image_id": "ImageId",
	"instance_type": "InstanceType",
	"monitoring": "Monitoring",
	"hibernation_options": "HibernationOptions",
	"license_specifications": "LicenseSpecifications",
	"instance_initiated_shutdown_behavior": "InstanceInitiatedShutdownBehavior",
	"cpu_options": "CpuOptions",
	"security_group_ids": "SecurityGroupIds",
	"key_name": "KeyName",
	"disable_api_termination": "DisableApiTermination",
	"instance_market_options": "InstanceMarketOptions",
	"ram_disk_id": "RamDiskId",
	"capacity_reservation_specification": "CapacityReservationSpecification",
	"credit_specification": "CreditSpecification",
}

func propertyLaunchTemplateLaunchTemplateData(extras...string) *schema.Resource {
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
			"security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateTagSpecification(),
				Optional: true,
			},
			"user_data": {
				Type: schema.TypeString,
				Optional: true,
			},
			"block_device_mappings": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateBlockDeviceMapping(),
				Optional: true,
			},
			"iam_instance_profile": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateIamInstanceProfile(),
				Optional: true,
				MaxItems: 1,
			},
			"kernel_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"elastic_gpu_specifications": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateElasticGpuSpecification(),
				Optional: true,
			},
			"elastic_inference_accelerators": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateLaunchTemplateElasticInferenceAccelerator(),
				Optional: true,
			},
			"placement": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplatePlacement(),
				Optional: true,
				MaxItems: 1,
			},
			"network_interfaces": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateNetworkInterface(),
				Optional: true,
			},
			"image_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"monitoring": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateMonitoring(),
				Optional: true,
				MaxItems: 1,
			},
			"hibernation_options": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateHibernationOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"license_specifications": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateLicenseSpecification(),
				Optional: true,
			},
			"instance_initiated_shutdown_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cpu_options": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateCpuOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"key_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"disable_api_termination": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"instance_market_options": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateInstanceMarketOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"ram_disk_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"capacity_reservation_specification": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateCapacityReservationSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"credit_specification": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateCreditSpecification(),
				Optional: true,
				MaxItems: 1,
			},
		},
	}
}
