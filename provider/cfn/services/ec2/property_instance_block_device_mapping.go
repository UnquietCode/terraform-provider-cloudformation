// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ec2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyInstanceBlockDeviceMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"ebs": {
				Type: schema.TypeList,
				Elem: propertyInstanceEbs(),
				Required: false,
				MaxItems: 1,
			},
			"no_device": {
				Type: schema.TypeList,
				Elem: propertyInstanceNoDevice(),
				Required: false,
				MaxItems: 1,
			},
			"virtual_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}