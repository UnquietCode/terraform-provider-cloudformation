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

func propertyVolume() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"docker_volume_configuration": {
				Type: schema.TypeList,
				Elem: propertyDockerVolumeConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"host": {
				Type: schema.TypeList,
				Elem: propertyHostVolumeProperties(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
		},
	}
}