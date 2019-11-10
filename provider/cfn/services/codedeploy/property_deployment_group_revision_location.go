// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package codedeploy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeploymentGroupRevisionLocation() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"git_hub_location": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupGitHubLocation(),
				Required: false,
				MaxItems: 1,
			},
			"revision_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"s3_location": {
				Type: schema.TypeList,
				Elem: propertyDeploymentGroupS3Location(),
				Required: false,
				MaxItems: 1,
			},
		},
	}
}