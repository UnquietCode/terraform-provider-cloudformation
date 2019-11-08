// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyProvisioningPreferences() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"stack_set_accounts": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"stack_set_failure_tolerance_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"stack_set_max_concurrency_percentage": {
				Type: schema.TypeInt,
				Required: false,
			},
			"stack_set_max_concurrency_count": {
				Type: schema.TypeInt,
				Required: false,
			},
			"stack_set_regions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"stack_set_operation_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"stack_set_failure_tolerance_percentage": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}