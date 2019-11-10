// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTaskDefinitionDevice() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"container_path": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"host_path": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"permissions": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
		},
	}
}