// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-environment.html

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
		return &schema.Resource{ Schema: map[string]*schema.Schema{} }
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
				Optional: true,
			},
			"privileged_mode": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"image_pull_credentials_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"image": {
				Type: schema.TypeString,
				Required: true,
			},
			"registry_credential": {
				Type: schema.TypeList,
				Elem: propertyProjectRegistryCredential(),
				Optional: true,
				MaxItems: 1,
			},
			"compute_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"certificate": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}