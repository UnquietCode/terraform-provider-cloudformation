// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ecs

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyTaskDefinitionLinuxParameters(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capabilities": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionKernelCapabilities(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"devices": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionDevice(),
				Required: false,
				ForceNew: true,
			},
			"init_process_enabled": {
				Type: schema.TypeBool,
				Required: false,
				ForceNew: true,
			},
			"shared_memory_size": {
				Type: schema.TypeInt,
				Required: false,
				ForceNew: true,
			},
			"tmpfs": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionTmpfs(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}