// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package ssm

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyPatchBaselineRule(extras...string) *schema.Resource {
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
			"enable_non_security": {
				Type: schema.TypeBool,
				Required: false,
			},
			"patch_filter_group": {
				Type: schema.TypeList,
				Elem: propertyPatchBaselinePatchFilterGroup(),
				Required: false,
				MaxItems: 1,
			},
			"approve_after_days": {
				Type: schema.TypeInt,
				Required: false,
			},
			"compliance_level": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}