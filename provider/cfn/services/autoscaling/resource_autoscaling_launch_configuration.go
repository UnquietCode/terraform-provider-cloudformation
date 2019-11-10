// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceAutoScalingLaunchConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutoScalingLaunchConfigurationCreate,
		Read:   resourceAutoScalingLaunchConfigurationRead,
		Delete: resourceAutoScalingLaunchConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"block_device_mappings": {
				Type: schema.TypeSet,
				Elem: propertyLaunchConfigurationBlockDeviceMapping(),
				Optional: true,
				ForceNew: true,
			},
			"classic_link_vpc_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"classic_link_vpc_security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
				Set: schema.HashString,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"iam_instance_profile": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"image_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"instance_monitoring": {
				Type: schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kernel_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"key_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"launch_configuration_name": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"placement_tenancy": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"ram_disk_id": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				ForceNew: true,
				Set: schema.HashString,
			},
			"spot_price": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"user_data": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAutoScalingLaunchConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::LaunchConfiguration", ResourceAutoScalingLaunchConfiguration(), data, meta)
}

func resourceAutoScalingLaunchConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::LaunchConfiguration", ResourceAutoScalingLaunchConfiguration(), data, meta)
}

func resourceAutoScalingLaunchConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::LaunchConfiguration", ResourceAutoScalingLaunchConfiguration(), data, meta)
}

func resourceAutoScalingLaunchConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::LaunchConfiguration", data, meta)
}