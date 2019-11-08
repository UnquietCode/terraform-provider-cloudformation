// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package opsworks

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyBlockDeviceMapping() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"device_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"ebs": {
				Type: schema.TypeList,
				Elem: propertyEbsBlockDevice(),
				Required: false,
				MaxItems: 1,
			},
			"no_device": {
				Type: schema.TypeString,
				Required: false,
			},
			"virtual_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}