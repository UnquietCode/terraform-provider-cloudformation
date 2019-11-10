// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceLaunchConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceLaunchConfigurationCreate,
		Read:   resourceLaunchConfigurationRead,
		Update: resourceLaunchConfigurationUpdate,
		Delete: resourceLaunchConfigurationDelete,

		Schema: map[string]*schema.Schema{
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"block_device_mappings": {
				Type: schema.TypeSet,
				Elem: propertyLaunchConfigurationBlockDeviceMapping(),
				Required: false,
				ForceNew: true,
			},
			"classic_link_vpc_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"classic_link_vpc_security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"iam_instance_profile": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"image_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"instance_monitoring": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"instance_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kernel_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"key_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"launch_configuration_name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"placement_tenancy": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"ram_disk_id": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"spot_price": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"user_data": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}

func resourceLaunchConfigurationCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::AutoScaling::LaunchConfiguration", data, meta)
}

func resourceLaunchConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::AutoScaling::LaunchConfiguration", data, meta)
}

func resourceLaunchConfigurationUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::AutoScaling::LaunchConfiguration", data, meta)
}

func resourceLaunchConfigurationDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::AutoScaling::LaunchConfiguration", data, meta)
}