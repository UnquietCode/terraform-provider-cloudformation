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

func ResourceRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouteCreate,
		Read:   resourceRouteRead,
		Update: resourceRouteUpdate,
		Delete: resourceRouteDelete,

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

func resourceRouteCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Route", data, meta)
}

func resourceRouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Route", data, meta)
}

func resourceRouteUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Route", data, meta)
}

func resourceRouteDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Route", data, meta)
}