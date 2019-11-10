// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package emr

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyStepHadoopJarStepConfig(extras...string) *schema.Resource {
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
			"args": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				ForceNew: true,
				Set: schema.HashString,
			},
			"jar": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"main_class": {
				Type: schema.TypeString,
				Required: false,
				ForceNew: true,
			},
			"step_properties": {
				Type: schema.TypeSet,
				Elem: propertyStepKeyValue(),
				Required: false,
				ForceNew: true,
			},
		},
	}
}