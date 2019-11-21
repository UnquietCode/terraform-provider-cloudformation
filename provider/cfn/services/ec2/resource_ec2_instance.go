// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html

package ec2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const eC2InstanceType string = "AWS::EC2::Instance"

func ResourceEC2Instance() *schema.Resource {
	return &schema.Resource{
		Read: resourceEC2InstanceRead,
		Create: resourceEC2InstanceCreate,
		Update: resourceEC2InstanceUpdate,
		Delete: resourceEC2InstanceDelete,
		CustomizeDiff: resourceEC2InstanceCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"additional_info": {
				Type: schema.TypeString,
				Optional: true,
			},
			"affinity": {
				Type: schema.TypeString,
				Optional: true,
			},
			"availability_zone": {
				Type: schema.TypeString,
				Optional: true,
			},
			"block_device_mappings": {
				Type: schema.TypeList,
				Elem: propertyInstanceBlockDeviceMapping(),
				Optional: true,
			},
			"cpu_options": {
				Type: schema.TypeSet,
				Elem: propertyInstanceCpuOptions(),
				Optional: true,
				MaxItems: 1,
			},
			"credit_specification": {
				Type: schema.TypeSet,
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
			},
			"elastic_inference_accelerators": {
				Type: schema.TypeSet,
				Elem: propertyInstanceElasticInferenceAccelerator(),
				Optional: true,
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
			},
			"ipv6_addresses": {
				Type: schema.TypeList,
				Elem: propertyInstanceInstanceIpv6Address(),
				Optional: true,
			},
			"kernel_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"key_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"launch_template": {
				Type: schema.TypeSet,
				Elem: propertyInstanceLaunchTemplateSpecification(),
				Optional: true,
				MaxItems: 1,
			},
			"license_specifications": {
				Type: schema.TypeSet,
				Elem: propertyInstanceLicenseSpecification(),
				Optional: true,
			},
			"monitoring": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"network_interfaces": {
				Type: schema.TypeList,
				Elem: propertyInstanceNetworkInterface(),
				Optional: true,
			},
			"placement_group_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"private_ip_address": {
				Type: schema.TypeString,
				Optional: true,
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
			},
			"tags": misc.PropertyTags(),
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
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceEC2InstanceRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(eC2InstanceType, ResourceEC2Instance(), data, meta)
}

func resourceEC2InstanceCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(eC2InstanceType, ResourceEC2Instance(), data, meta)
}

func resourceEC2InstanceUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(eC2InstanceType, ResourceEC2Instance(), data, meta)
}

func resourceEC2InstanceDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(eC2InstanceType, data, meta)
}

func resourceEC2InstanceCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(eC2InstanceType, data, meta)
}
