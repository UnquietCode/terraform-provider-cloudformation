// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package kinesisanalyticsv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyApplicationMappingParameters() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"json_mapping_parameters": {
				Type: schema.TypeList,
				Elem: propertyApplicationJSONMappingParameters(),
				Required: false,
				MaxItems: 1,
			},
			"csv_mapping_parameters": {
				Type: schema.TypeList,
				Elem: propertyApplicationCSVMappingParameters(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}