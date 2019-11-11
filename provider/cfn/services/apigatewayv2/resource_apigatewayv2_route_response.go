// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-routeresponse.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayV2RouteResponse() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayV2RouteResponseExists,
		Read: resourceApiGatewayV2RouteResponseRead,
		Create: resourceApiGatewayV2RouteResponseCreate,
		Update: resourceApiGatewayV2RouteResponseUpdate,
		Delete: resourceApiGatewayV2RouteResponseDelete,
		CustomizeDiff: resourceApiGatewayV2RouteResponseCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"route_response_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"response_parameters": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"route_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"model_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"response_models": {
				Type: schema.TypeMap,
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

func resourceApiGatewayV2RouteResponseExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayV2RouteResponseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::RouteResponse", ResourceApiGatewayV2RouteResponse(), data, meta)
}

func resourceApiGatewayV2RouteResponseCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::RouteResponse", ResourceApiGatewayV2RouteResponse(), data, meta)
}

func resourceApiGatewayV2RouteResponseUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::RouteResponse", ResourceApiGatewayV2RouteResponse(), data, meta)
}

func resourceApiGatewayV2RouteResponseDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::RouteResponse", data, meta)
}

func resourceApiGatewayV2RouteResponseCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::ApiGatewayV2::RouteResponse", data, meta)
}

