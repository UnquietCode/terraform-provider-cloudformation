// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ecs

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTaskDefinitionHealthCheck() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
				ForceNew: true,
				Set: schema.HashString,
			},
			"interval": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"retries": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"start_period": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"timeout": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
		},
	}
}