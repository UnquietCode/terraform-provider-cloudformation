// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyStageCanarySetting() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"deployment_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"percent_traffic": {
				Type: schema.TypeFloat,
				Required: false,
			},
			"stage_variable_overrides": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"use_stage_cache": {
				Type: schema.TypeBool,
				Required: false,
			},
		},
	}
}