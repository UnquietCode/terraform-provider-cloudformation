// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyRevisionLocation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"git_hub_location": {
				Type: schema.TypeList,
				Elem: propertyGitHubLocation(),
				Required: false,
				MaxItems: 1,
			},
			"revision_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"s3_location": {
				Type: schema.TypeList,
				Elem: propertyS3Location(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}