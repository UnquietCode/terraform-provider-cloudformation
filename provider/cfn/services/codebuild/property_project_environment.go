// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codebuild

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyProjectEnvironment(extras...string) *schema.Resource {
	var count int64 = 0
	
	if len(extras) > 0 {
		if i, err := strconv.ParseInt(extras[0], 10, 32); err == nil {
			count = i
		}
	}
	
	if count >= 5 {
		return nil
	}
	
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"environment_variables": {
				Type: schema.TypeList,
				Elem: propertyProjectEnvironmentVariable(),
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
				Elem: propertyProjectRegistryCredential(),
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