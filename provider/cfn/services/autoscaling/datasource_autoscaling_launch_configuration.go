// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html

package autoscaling

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const autoScalingLaunchConfigurationType string = "AWS::AutoScaling::LaunchConfiguration"

func DatasourceAutoScalingLaunchConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: datasourceAutoScalingLaunchConfigurationRead,
		
		Schema: map[string]*schema.Schema{
			"associate_public_ip_address": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"block_device_mappings": {
				Type: schema.TypeSet,
				Elem: propertyLaunchConfigurationBlockDeviceMapping(),
				Optional: true,
			},
			"classic_link_vpc_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"classic_link_vpc_security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"ebs_optimized": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"iam_instance_profile": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"instance_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"instance_monitoring": {
				Type: schema.TypeBool,
				Optional: true,
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
			"launch_configuration_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"placement_tenancy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"ram_disk_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"security_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"spot_price": {
				Type: schema.TypeString,
				Optional: true,
			},
			"user_data": {
				Type: schema.TypeString,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceAutoScalingLaunchConfigurationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(autoScalingLaunchConfigurationType, DatasourceAutoScalingLaunchConfiguration(), data, meta)
}
