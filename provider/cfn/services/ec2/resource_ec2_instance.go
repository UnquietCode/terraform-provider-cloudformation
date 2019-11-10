// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceEC2Instance() *schema.Resource {
	return &schema.Resource{
		Create: resourceEC2InstanceCreate,
		Read:   resourceEC2InstanceRead,
		Update: resourceEC2InstanceUpdate,
		Delete: resourceEC2InstanceDelete,

		Schema: map[string]*schema.Schema{
			"additional_info": {
				Type: schema.TypeString,
				Required: false,
			},
			"affinity": {
				Type: schema.TypeString,
				Required: false,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"block_device_mappings": {
				Type: schema.TypeList,
				Elem: propertyInstanceBlockDeviceMapping(),
				Required: false,
			},
			"cpu_options": {
				Type: schema.TypeList,
				Elem: propertyInstanceCpuOptions(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"credit_specification": {
				Type: schema.TypeList,
				Elem: propertyInstanceCreditSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"disable_api_termination": {
				Type: schema.TypeBool,
				Required: false,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Required: false,
			},
			"elastic_gpu_specifications": {
				Type: schema.TypeSet,
				Elem: propertyInstanceElasticGpuSpecification(),
				Required: false,
				ForceNew: true,
			},
			"elastic_inference_accelerators": {
				Type: schema.TypeSet,
				Elem: propertyInstanceElasticInferenceAccelerator(),
				Required: false,
				ForceNew: true,
			},
			"host_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"iam_instance_profile": {
				Type: schema.TypeString,
				Required: false,
			},
			"image_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"instance_initiated_shutdown_behavior": {
				Type: schema.TypeString,
				Required: false,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"ipv6_addresses": {
				Type: schema.TypeList,
				Elem: propertyInstanceInstanceIpv6Address(),
				Required: false,
				ForceNew: true,
			},
			"kernel_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"key_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"launch_template": {
				Type: schema.TypeList,
				Elem: propertyInstanceLaunchTemplateSpecification(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"license_specifications": {
				Type: schema.TypeSet,
				Elem: propertyInstanceLicenseSpecification(),
				Required: false,
				ForceNew: true,
			},
			"monitoring": {
				Type: schema.TypeBool,
				Required: false,
			},
			"network_interfaces": {
				Type: schema.TypeList,
				Elem: propertyInstanceNetworkInterface(),
				Required: false,
				ForceNew: true,
			},
			"placement_group_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ramdisk_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"source_dest_check": {
				Type: schema.TypeBool,
				Required: false,
			},
			"ssm_associations": {
				Type: schema.TypeList,
				Elem: propertyInstanceSsmAssociation(),
				Required: false,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Required: false,
			},
			"tenancy": {
				Type: schema.TypeString,
				Required: false,
			},
			"user_data": {
				Type: schema.TypeString,
				Required: false,
			},
			"volumes": {
				Type: schema.TypeList,
				Elem: propertyInstanceVolume(),
				Required: false,
			},
		},
	}
}

func resourceEC2InstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::EC2::Instance", data, meta)
}

func resourceEC2InstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::EC2::Instance", data, meta)
}

func resourceEC2InstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::EC2::Instance", data, meta)
}

func resourceEC2InstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::EC2::Instance", data, meta)
}