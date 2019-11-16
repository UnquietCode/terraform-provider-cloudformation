// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-mapping.html

package autoscaling

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var launchConfigurationBlockDeviceMappingProperties map[string]string = map[string]string{
	"device_name": "DeviceName",
	"ebs": "Ebs",
	"no_device": "NoDevice",
	"virtual_name": "VirtualName",
}

func propertyLaunchConfigurationBlockDeviceMapping(extras...string) *schema.Resource {
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
			"device_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ebs": {
				Type: schema.TypeList,
				Elem: propertyLaunchConfigurationBlockDevice(),
				Optional: true,
				MaxItems: 1,
			},
			"no_device": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"virtual_name": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
