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

func propertySpotFleetLaunchSpecification() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"block_device_mappings": {
				Type: schema.TypeSet,
				Elem: propertyBlockDeviceMapping(),
				Required: false,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Required: false,
			},
			"iam_instance_profile": {
				Type: schema.TypeList,
				Elem: propertyIamInstanceProfileSpecification(),
				Required: false,
				MaxItems: 1,
			},
			"image_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"kernel_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"key_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"monitoring": {
				Type: schema.TypeList,
				Elem: propertySpotFleetMonitoring(),
				Required: false,
				MaxItems: 1,
			},
			"network_interfaces": {
				Type: schema.TypeSet,
				Elem: propertyInstanceNetworkInterfaceSpecification(),
				Required: false,
			},
			"placement": {
				Type: schema.TypeList,
				Elem: propertySpotPlacement(),
				Required: false,
				MaxItems: 1,
			},
			"ramdisk_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: propertyGroupIdentifier(),
				Required: false,
			},
			"spot_price": {
				Type: schema.TypeString,
				Required: false,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"tag_specifications": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetTagSpecification(),
				Required: false,
			},
			"user_data": {
				Type: schema.TypeString,
				Required: false,
			},
			"weighted_capacity": {
				Type: schema.TypeFloat,
				Required: false,
			},
		},
	}
}