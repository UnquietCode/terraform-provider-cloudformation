// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html

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
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"block_device_mappings": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetBlockDeviceMapping(),
				Optional: true,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"iam_instance_profile": {
				Type: schema.TypeList,
				Elem: propertySpotFleetIamInstanceProfileSpecification(),
				Optional: true,
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
				Optional: true,
			},
			"key_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"monitoring": {
				Type: schema.TypeList,
				Elem: propertySpotFleetSpotFleetMonitoring(),
				Optional: true,
				MaxItems: 1,
			},
			"network_interfaces": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetInstanceNetworkInterfaceSpecification(),
				Optional: true,
			},
			"placement": {
				Type: schema.TypeList,
				Elem: propertySpotFleetSpotPlacement(),
				Optional: true,
				MaxItems: 1,
			},
			"ramdisk_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetGroupIdentifier(),
				Optional: true,
			},
			"spot_price": {
				Type: schema.TypeString,
				Optional: true,
			},
			"subnet_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tag_specifications": {
				Type: schema.TypeSet,
				Elem: propertySpotFleetSpotFleetTagSpecification(),
				Optional: true,
			},
			"user_data": {
				Type: schema.TypeString,
				Optional: true,
			},
			"weighted_capacity": {
				Type: schema.TypeFloat,
				Optional: true,
			},
		},
	}
}

