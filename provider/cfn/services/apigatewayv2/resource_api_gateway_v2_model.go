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

func ResourceApiGatewayV2Model() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayV2ModelCreate,
		Read:   resourceApiGatewayV2ModelRead,
		Update: resourceApiGatewayV2ModelUpdate,
		Delete: resourceApiGatewayV2ModelDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"content_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"schema": {
				Type: schema.TypeMap,
				Required: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApiGatewayV2ModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Model", data, meta)
}

func resourceApiGatewayV2ModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Model", data, meta)
}

func resourceApiGatewayV2ModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Model", data, meta)
}

func resourceApiGatewayV2ModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Model", data, meta)
}