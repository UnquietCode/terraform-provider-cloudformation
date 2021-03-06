// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions.html

package codepipeline

import (
	"strconv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func propertyPipelineActionDeclaration(extras...string) *schema.Resource {
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
			"action_type_id": {
				Type: schema.TypeList,
				Elem: propertyPipelineActionTypeId(),
				Required: true,
				MaxItems: 1,
			},
			"configuration": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"input_artifacts": {
				Type: schema.TypeSet,
				Elem: propertyPipelineInputArtifact(),
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type: schema.TypeString,
				Optional: true,
			},
			"output_artifacts": {
				Type: schema.TypeSet,
				Elem: propertyPipelineOutputArtifact(),
				Optional: true,
			},
			"region": {
				Type: schema.TypeString,
				Optional: true,
			},
			"role_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"run_order": {
				Type: schema.TypeInt,
				Optional: true,
			},
		},
	}
}
