// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTaskDefinitionVolume() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"docker_volume_configuration": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionDockerVolumeConfiguration(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"host": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionHostVolumeProperties(),
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