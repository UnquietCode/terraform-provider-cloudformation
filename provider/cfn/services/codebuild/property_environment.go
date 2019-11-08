// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codebuild

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyEnvironment() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"environment_variables": {
				Type: schema.TypeList,
				Elem: propertyEnvironmentVariable(),
				Required: false,
			},
			"privileged_mode": {
				Type: schema.TypeBool,
				Required: false,
			},
			"image_pull_credentials_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"image": {
				Type: schema.TypeString,
				Required: true,
			},
			"registry_credential": {
				Type: schema.TypeList,
				Elem: propertyRegistryCredential(),
				Required: false,
				MaxItems: 1,
			},
			"compute_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"certificate": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}