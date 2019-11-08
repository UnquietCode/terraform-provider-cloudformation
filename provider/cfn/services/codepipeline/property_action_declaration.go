// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyActionDeclaration() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"action_type_id": {
				Type: schema.TypeList,
				Elem: propertyActionTypeId(),
				Required: true,
				MaxItems: 1,
			},
			"configuration": {
				Type: schema.TypeMap,
				Required: false,
			},
			"input_artifacts": {
				Type: schema.TypeSet,
				Elem: propertyInputArtifact(),
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"output_artifacts": {
				Type: schema.TypeSet,
				Elem: propertyOutputArtifact(),
				Required: false,
			},
			"region": {
				Type: schema.TypeString,
				Required: false,
			},
			"role_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"run_order": {
				Type: schema.TypeInt,
				Required: false,
			},
		},
	}
}