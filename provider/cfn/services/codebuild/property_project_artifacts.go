// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyProjectArtifacts() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": {
				Type: schema.TypeString,
				Required: false,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"artifact_identifier": {
				Type: schema.TypeString,
				Required: false,
			},
			"override_artifact_name": {
				Type: schema.TypeBool,
				Required: false,
			},
			"packaging": {
				Type: schema.TypeString,
				Required: false,
			},
			"encryption_disabled": {
				Type: schema.TypeBool,
				Required: false,
			},
			"location": {
				Type: schema.TypeString,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
			"namespace_type": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}