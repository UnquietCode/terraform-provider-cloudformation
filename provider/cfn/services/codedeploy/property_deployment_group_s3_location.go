// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-s3location.html

package codedeploy

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func propertyDeploymentGroupS3Location(extras...string) *schema.Resource {
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
			"bucket": {
				Type: schema.TypeString,
				Required: true,
			},
			"bundle_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"e_tag": {
				Type: schema.TypeString,
				Optional: true,
			},
			"key": {
				Type: schema.TypeString,
				Required: true,
			},
			"version": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}
