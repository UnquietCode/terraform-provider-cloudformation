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

func propertySpotFleetSpotFleetLaunchSpecification(extras...string) *schema.Resource {
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
			"block_device_mappings": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetBlockDeviceMapping(),
				Required: false,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Required: false,
			},
			"iam_instance_profile": {
				Type: schema.TypeList,
				Elem: propertySpotFleetIamInstanceProfileSpecification(),
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
				Elem: propertySpotFleetSpotFleetMonitoring(),
				Required: false,
				MaxItems: 1,
			},
			"network_interfaces": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetInstanceNetworkInterfaceSpecification(),
				Required: false,
			},
			"placement": {
				Type: schema.TypeList,
				Elem: propertySpotFleetSpotPlacement(),
				Required: false,
				MaxItems: 1,
			},
			"ramdisk_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetGroupIdentifier(),
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
				Elem: propertySpotFleetSpotFleetTagSpecification(),
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