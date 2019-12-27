// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integration.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2IntegrationType string = "AWS::ApiGatewayV2::Integration"

func DatasourceApiGatewayV2Integration() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayV2IntegrationRead,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"template_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"connection_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"integration_method": {
				Type: schema.TypeString,
				Optional: true,
			},
			"passthrough_behavior": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_parameters": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"integration_uri": {
				Type: schema.TypeString,
				Optional: true,
			},
			"payload_format_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"credentials_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_templates": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"timeout_in_millis": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"content_handling_strategy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"integration_type": {
				Type: schema.TypeString,
				Required: true,
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

func datasourceApiGatewayV2IntegrationRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2IntegrationType, DatasourceApiGatewayV2Integration(), data, meta)
}
