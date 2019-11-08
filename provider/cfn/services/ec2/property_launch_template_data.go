// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLaunchTemplateData() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"tag_specifications": {
				Type: schema.TypeList,
				Elem: propertyTagSpecification(),
				Required: false,
			},
			"user_data": {
				Type: schema.TypeString,
				Required: false,
			},
			"block_device_mappings": {
				Type: schema.TypeList,
				Elem: propertyBlockDeviceMapping(),
				Required: false,
			},
			"iam_instance_profile": {
				Type: schema.TypeList,
				Elem: propertyIamInstanceProfile(),
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
				Elem: propertyElasticGpuSpecification(),
				Required: false,
			},
			"elastic_inference_accelerators": {
				Type: schema.TypeList,
				Elem: propertyLaunchTemplateElasticInferenceAccelerator(),
				Required: false,
			},
			"placement": {
				Type: schema.TypeList,
				Elem: propertyPlacement(),
				Required: false,
				MaxItems: 1,
			},
			"network_interfaces": {
				Type: schema.TypeList,
				Elem: propertyNetworkInterface(),
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
				Elem: propertyMonitoring(),
				Required: false,
				MaxItems: 1,
			},
			"hibernation_options": {
				Type: schema.TypeList,
				Elem: propertyHibernationOptions(),
				Required: false,
				MaxItems: 1,
			},
			"license_specifications": {
				Type: schema.TypeList,
				Elem: propertyLicenseSpecification(),
				Required: false,
			},
			"instance_initiated_shutdown_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"cpu_options": {
				Type: schema.TypeList,
				Elem: propertyCpuOptions(),
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
				Elem: propertyInstanceMarketOptions(),
				Required: false,
				MaxItems: 1,
			},
			"ram_disk_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"capacity_reservation_specification": {
				Type: schema.TypeList,
				Elem: propertyCapacityReservationSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"credit_specification": {
				Type: schema.TypeList,
				Elem: propertyCreditSpecification(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}