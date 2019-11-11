// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html

package codepipeline

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceCodePipelineWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceCodePipelineWebhookCreate,
		Read:   resourceCodePipelineWebhookRead,
		Update: resourceCodePipelineWebhookUpdate,
		Delete: resourceCodePipelineWebhookDelete,

		Schema: map[string]*schema.Schema{
			"url": {
				Type: schema.TypeString,
				Computed: true,
			},
			"authentication_configuration": {
				Type: schema.TypeList,
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
				ForceNew: true,
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

func resourceCodePipelineWebhookCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::CodePipeline::Webhook", ResourceCodePipelineWebhook(), data, meta)
}

func resourceCodePipelineWebhookRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::CodePipeline::Webhook", ResourceCodePipelineWebhook(), data, meta)
}

func resourceCodePipelineWebhookUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::CodePipeline::Webhook", ResourceCodePipelineWebhook(), data, meta)
}

func resourceCodePipelineWebhookDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::CodePipeline::Webhook", data, meta)
}