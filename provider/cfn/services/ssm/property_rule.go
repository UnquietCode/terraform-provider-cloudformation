// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package ssm

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRule() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_non_security": {
				Type: schema.TypeBool,
				Required: false,
			},
			"patch_filter_group": {
				Type: schema.TypeList,
				Elem: propertyPatchFilterGroup(),
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