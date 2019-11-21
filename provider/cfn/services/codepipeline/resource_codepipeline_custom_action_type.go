// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html

package codepipeline

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codePipelineCustomActionTypeType string = "AWS::CodePipeline::CustomActionType"

func ResourceCodePipelineCustomActionType() *schema.Resource {
	return &schema.Resource{
		Read: resourceCodePipelineCustomActionTypeRead,
		Create: resourceCodePipelineCustomActionTypeCreate,
		Update: resourceCodePipelineCustomActionTypeUpdate,
		Delete: resourceCodePipelineCustomActionTypeDelete,
		CustomizeDiff: resourceCodePipelineCustomActionTypeCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"category": {
				Type: schema.TypeString,
				Required: true,
			},
			"configuration_properties": {
				Type: schema.TypeSet,
				Elem: propertyCustomActionTypeConfigurationProperties(),
				Optional: true,
			},
			"input_artifact_details": {
				Type: schema.TypeSet,
				Elem: propertyCustomActionTypeArtifactDetails(),
				Required: true,
				MaxItems: 1,
			},
			"output_artifact_details": {
				Type: schema.TypeSet,
				Elem: propertyCustomActionTypeArtifactDetails(),
				Required: true,
				MaxItems: 1,
			},
			"the_provider": {
				Type: schema.TypeString,
				Required: true,
			},
			"settings": {
				Type: schema.TypeSet,
				Elem: propertyCustomActionTypeSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"version": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceCodePipelineCustomActionTypeRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codePipelineCustomActionTypeType, ResourceCodePipelineCustomActionType(), data, meta)
}

func resourceCodePipelineCustomActionTypeCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codePipelineCustomActionTypeType, ResourceCodePipelineCustomActionType(), data, meta)
}

func resourceCodePipelineCustomActionTypeUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codePipelineCustomActionTypeType, ResourceCodePipelineCustomActionType(), data, meta)
}

func resourceCodePipelineCustomActionTypeDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codePipelineCustomActionTypeType, data, meta)
}

func resourceCodePipelineCustomActionTypeCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codePipelineCustomActionTypeType, data, meta)
}
