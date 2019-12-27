// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html

package codepipeline

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const codePipelineWebhookType string = "AWS::CodePipeline::Webhook"

func DatasourceCodePipelineWebhook() *schema.Resource {
	return &schema.Resource{
		Read: datasourceCodePipelineWebhookRead,
		
		Schema: map[string]*schema.Schema{
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceCodePipelineWebhookRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(codePipelineWebhookType, DatasourceCodePipelineWebhook(), data, meta)
}
