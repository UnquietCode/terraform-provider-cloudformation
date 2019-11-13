// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
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
		Exists: resourceApiGatewayV2RouteExists,
		Read:   resourceApiGatewayV2RouteRead,
		Create: resourceApiGatewayV2RouteCreate,
		Update: resourceApiGatewayV2RouteUpdate,
		Delete: resourceApiGatewayV2RouteDelete,
		CustomizeDiff: resourceApiGatewayV2RouteCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"target": {
				Type: schema.TypeString,
				Optional: true,
			},
			"route_response_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_models": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"operation_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorization_scopes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"api_key_required": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"route_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorization_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"model_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"request_parameters": {
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

func resourceApiGatewayV2RouteExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayV2RouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Route", ResourceApiGatewayV2Route(), data, meta)
}

func resourceApiGatewayV2RouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Route", ResourceApiGatewayV2Route(), data, meta)
}

func resourceApiGatewayV2RouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Route", ResourceApiGatewayV2Route(), data, meta)
}

func resourceApiGatewayV2RouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Route", data, meta)
}

func resourceApiGatewayV2RouteCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}