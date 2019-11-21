// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2IntegrationResponseType string = "AWS::ApiGatewayV2::IntegrationResponse"

func ResourceApiGatewayV2IntegrationResponse() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayV2IntegrationResponseRead,
		Create: resourceApiGatewayV2IntegrationResponseCreate,
		Update: resourceApiGatewayV2IntegrationResponseUpdate,
		Delete: resourceApiGatewayV2IntegrationResponseDelete,
		CustomizeDiff: resourceApiGatewayV2IntegrationResponseCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"response_templates": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"template_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"response_parameters": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"content_handling_strategy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"integration_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"integration_response_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_id": {
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

func resourceApiGatewayV2IntegrationResponseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2IntegrationResponseType, ResourceApiGatewayV2IntegrationResponse(), data, meta)
}

func resourceApiGatewayV2IntegrationResponseCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayV2IntegrationResponseType, ResourceApiGatewayV2IntegrationResponse(), data, meta)
}

func resourceApiGatewayV2IntegrationResponseUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayV2IntegrationResponseType, ResourceApiGatewayV2IntegrationResponse(), data, meta)
}

func resourceApiGatewayV2IntegrationResponseDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayV2IntegrationResponseType, data, meta)
}

func resourceApiGatewayV2IntegrationResponseCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayV2IntegrationResponseType, data, meta)
}
