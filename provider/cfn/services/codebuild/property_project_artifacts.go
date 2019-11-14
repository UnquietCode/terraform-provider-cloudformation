// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 14-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-artifacts.html

package codebuild

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyProjectArtifacts(extras...string) *schema.Resource {
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
			"path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"type": {
				Type: schema.TypeString,
				Required: true,
			},
			"artifact_identifier": {
				Type: schema.TypeString,
				Optional: true,
			},
			"override_artifact_name": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"packaging": {
				Type: schema.TypeString,
				Optional: true,
			},
			"encryption_disabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"location": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"namespace_type": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

