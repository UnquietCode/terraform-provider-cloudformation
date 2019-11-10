// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package autoscaling

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyLaunchConfigurationBlockDeviceMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ebs": {
				Type: schema.TypeList,
				Elem: propertyLaunchConfigurationBlockDevice(),
				Required: false,
				MaxItems: 1,
			},
			"no_device": {
				Type: schema.TypeBool,
				Required: false,
			},
			"virtual_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}