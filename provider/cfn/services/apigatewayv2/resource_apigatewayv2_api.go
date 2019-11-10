// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayV2Api() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayV2ApiCreate,
		Read:   resourceApiGatewayV2ApiRead,
		Update: resourceApiGatewayV2ApiUpdate,
		Delete: resourceApiGatewayV2ApiDelete,

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

func resourceApiGatewayV2ApiCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Api", data, meta)
}

func resourceApiGatewayV2ApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Api", data, meta)
}

func resourceApiGatewayV2ApiUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Api", data, meta)
}

func resourceApiGatewayV2ApiDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Api", data, meta)
}