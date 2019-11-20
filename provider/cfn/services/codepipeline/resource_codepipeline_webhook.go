// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codePipelineWebhookType string = "AWS::CodePipeline::Webhook"

func ResourceCodePipelineWebhook() *schema.Resource {
	return &schema.Resource{
		Read: resourceCodePipelineWebhookRead,
		Create: resourceCodePipelineWebhookCreate,
		Update: resourceCodePipelineWebhookUpdate,
		Delete: resourceCodePipelineWebhookDelete,
		CustomizeDiff: resourceCodePipelineWebhookCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"authentication_configuration": {
				Type: schema.TypeSet,
				Elem: propertyWebhookWebhookAuthConfiguration(),
				Required: true,
				MaxItems: 1,
			},
			"filters": {
				Type: schema.TypeList,
				Elem: propertyWebhookWebhookFilterRule(),
				Required: true,
			},
			"authentication": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_pipeline": {
				Type: schema.TypeString,
				Required: true,
			},
			"target_action": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target_pipeline_version": {
				Type: schema.TypeInt,
				Required: true,
			},
			"register_with_third_party": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCodePipelineWebhookRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codePipelineWebhookType, ResourceCodePipelineWebhook(), data, meta)
}

func resourceCodePipelineWebhookCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(codePipelineWebhookType, ResourceCodePipelineWebhook(), data, meta)
}

func resourceCodePipelineWebhookUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(codePipelineWebhookType, ResourceCodePipelineWebhook(), data, meta)
}

func resourceCodePipelineWebhookDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(codePipelineWebhookType, data, meta)
}

func resourceCodePipelineWebhookCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(codePipelineWebhookType, data, meta)
}
