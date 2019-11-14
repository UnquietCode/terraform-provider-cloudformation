// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicecatalog-cloudformationprovisionedproduct-provisioningpreferences.html

package servicecatalog

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCloudFormationProvisionedProductProvisioningPreferences(extras...string) *schema.Resource {
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
			"stack_set_accounts": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"stack_set_failure_tolerance_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"stack_set_max_concurrency_percentage": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"stack_set_max_concurrency_count": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"stack_set_regions": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"stack_set_operation_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"stack_set_failure_tolerance_percentage": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
