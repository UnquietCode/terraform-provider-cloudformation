// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayV2IntegrationResponse() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayV2IntegrationResponseCreate,
		Read:   resourceApiGatewayV2IntegrationResponseRead,
		Update: resourceApiGatewayV2IntegrationResponseUpdate,
		Delete: resourceApiGatewayV2IntegrationResponseDelete,

		Schema: map[string]*schema.Schema{
			"response_templates": {
				Type: schema.TypeMap,
				Required: false,
			},
			"template_selection_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"response_parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"content_handling_strategy": {
				Type: schema.TypeString,
				Required: false,
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
		},
	}
}

func resourceApiGatewayV2IntegrationResponseCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::IntegrationResponse", data, meta)
}

func resourceApiGatewayV2IntegrationResponseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::IntegrationResponse", data, meta)
}

func resourceApiGatewayV2IntegrationResponseUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::IntegrationResponse", data, meta)
}

func resourceApiGatewayV2IntegrationResponseDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::IntegrationResponse", data, meta)
}