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

func ResourceApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiCreate,
		Read:   resourceApiRead,
		Update: resourceApiUpdate,
		Delete: resourceApiDelete,

		Schema: map[string]*schema.Schema{
			"route_selection_expression": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"version": {
				Type: schema.TypeString,
				Required: false,
			},
			"protocol_type": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"disable_schema_validation": {
				Type: schema.TypeBool,
				Required: false,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
			"name": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_key_selection_expression": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceApiCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Api", data, meta)
}

func resourceApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Api", data, meta)
}

func resourceApiUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Api", data, meta)
}

func resourceApiDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Api", data, meta)
}