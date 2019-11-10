// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceRouteResponse() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouteResponseCreate,
		Read:   resourceRouteResponseRead,
		Update: resourceRouteResponseUpdate,
		Delete: resourceRouteResponseDelete,

		Schema: map[string]*schema.Schema{
			"route_response_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"response_parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
			"route_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"model_selection_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"response_models": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceRouteResponseCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::RouteResponse", data, meta)
}

func resourceRouteResponseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::RouteResponse", data, meta)
}

func resourceRouteResponseUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::RouteResponse", data, meta)
}

func resourceRouteResponseDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::RouteResponse", data, meta)
}