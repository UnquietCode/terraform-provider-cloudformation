// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDockerVolumeConfiguration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoprovision": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"driver": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"driver_opts": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"labels": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
			},
			"scope": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}