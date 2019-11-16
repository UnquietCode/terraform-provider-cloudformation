// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-blockdevicemapping.html

package ec2

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var launchTemplateBlockDeviceMappingProperties map[string]string = map[string]string{
	"ebs": "Ebs",
	"no_device": "NoDevice",
	"virtual_name": "VirtualName",
	"device_name": "DeviceName",
}

func propertyLaunchTemplateBlockDeviceMapping(extras...string) *schema.Resource {
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
			"ebs": {
				Type: schema.TypeSet,
				Elem: propertyLaunchTemplateEbs(),
				Optional: true,
				MaxItems: 1,
			},
			"no_device": {
				Type: schema.TypeString,
				Optional: true,
			},
			"virtual_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"device_name": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
