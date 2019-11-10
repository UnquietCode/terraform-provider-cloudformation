// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html

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
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"private_dns_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"private_ip": {
				Type: schema.TypeString,
				Computed: true,
			},
			"public_dns_name": {
				Type: schema.TypeString,
				Computed: true,
			},
			"public_ip": {
				Type: schema.TypeString,
				Computed: true,
			},
			"additional_info": {
				Type: schema.TypeString,
				Optional: true,
			},
			"affinity": {
				Type: schema.TypeString,
				Optional: true,
			},
			"block_device_mappings": {
				Type: schema.TypeList,
				Elem: propertyInstanceBlockDeviceMapping(),
				Optional: true,
			},
			"cpu_options": {
				Type: schema.TypeList,
				Elem: propertyInstanceCpuOptions(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"credit_specification": {
				Type: schema.TypeList,
				Elem: propertyInstanceCreditSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"disable_api_termination": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"elastic_gpu_specifications": {
				Type: schema.TypeSet,
				Elem: propertyInstanceElasticGpuSpecification(),
				Optional: true,
				ForceNew: true,
			},
			"elastic_inference_accelerators": {
				Type: schema.TypeSet,
				Elem: propertyInstanceElasticInferenceAccelerator(),
				Optional: true,
				ForceNew: true,
			},
			"host_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"iam_instance_profile": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"instance_initiated_shutdown_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ipv6_address_count": {
				Type: schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"ipv6_addresses": {
				Type: schema.TypeList,
				Elem: propertyInstanceInstanceIpv6Address(),
				Optional: true,
				ForceNew: true,
			},
			"kernel_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"key_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"launch_template": {
				Type: schema.TypeList,
				Elem: propertyInstanceLaunchTemplateSpecification(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"license_specifications": {
				Type: schema.TypeSet,
				Elem: propertyInstanceLicenseSpecification(),
				Optional: true,
				ForceNew: true,
			},
			"monitoring": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"network_interfaces": {
				Type: schema.TypeList,
				Elem: propertyInstanceNetworkInterface(),
				Optional: true,
				ForceNew: true,
			},
			"placement_group_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ramdisk_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_group_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"security_groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
			},
			"source_dest_check": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"ssm_associations": {
				Type: schema.TypeList,
				Elem: propertyInstanceSsmAssociation(),
				Optional: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"tags": {
				Type: schema.TypeList,
				Elem: misc.PropertyTag(),
				Optional: true,
			},
			"tenancy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_data": {
				Type: schema.TypeString,
				Optional: true,
			},
			"volumes": {
				Type: schema.TypeList,
				Elem: propertyInstanceVolume(),
				Optional: true,
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