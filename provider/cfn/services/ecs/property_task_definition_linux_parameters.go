// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-linuxparameters.html

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
	    return &schema.Resource{ Schema: map[string]*schema.Schema{} }
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capabilities": {
				Type: schema.TypeList,
				Elem: propertyTaskDefinitionKernelCapabilities(),
				Optional: true,
				MaxItems: 1,
			},
			"devices": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionDevice(),
				Optional: true,
			},
			"init_process_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"max_swap": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"shared_memory_size": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"swappiness": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"tmpfs": {
				Type: schema.TypeSet,
				Elem: propertyTaskDefinitionTmpfs(),
				Optional: true,
			},
		},
	}
}
