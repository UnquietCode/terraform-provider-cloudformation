// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package config

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyConfigurationRecorderRecordingGroup() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"all_supported": {
				Type: schema.TypeBool,
				Required: false,
			},
			"include_global_resource_types": {
				Type: schema.TypeBool,
				Required: false,
			},
			"resource_types": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
		},
	}
}