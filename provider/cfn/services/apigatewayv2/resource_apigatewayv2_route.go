// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayV2Route() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayV2RouteCreate,
		Read:   resourceApiGatewayV2RouteRead,
		Update: resourceApiGatewayV2RouteUpdate,
		Delete: resourceApiGatewayV2RouteDelete,

		Schema: map[string]*schema.Schema{
			"target": {
				Type: schema.TypeString,
				Required: false,
			},
			"route_response_selection_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"authorizer_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"request_models": {
				Type: schema.TypeMap,
				Required: false,
			},
			"operation_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"authorization_scopes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"api_key_required": {
				Type: schema.TypeBool,
				Required: false,
			},
			"route_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorization_type": {
				Type: schema.TypeString,
				Required: false,
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
			"request_parameters": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceApiGatewayV2RouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Route", data, meta)
}

func resourceApiGatewayV2RouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Route", data, meta)
}

func resourceApiGatewayV2RouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Route", data, meta)
}

func resourceApiGatewayV2RouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Route", data, meta)
}