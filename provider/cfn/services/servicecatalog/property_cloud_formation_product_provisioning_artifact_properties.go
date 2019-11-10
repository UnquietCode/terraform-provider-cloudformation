// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package servicecatalog

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyCloudFormationProductProvisioningArtifactProperties() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"disable_template_validation": {
				Type: schema.TypeBool,
				Required: false,
			},
			"info": {
				Type: schema.TypeMap,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}