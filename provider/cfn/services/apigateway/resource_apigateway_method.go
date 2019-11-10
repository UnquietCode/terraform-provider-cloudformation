// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayMethod() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayMethodCreate,
		Read:   resourceApiGatewayMethodRead,
		Update: resourceApiGatewayMethodUpdate,
		Delete: resourceApiGatewayMethodDelete,

		Schema: map[string]*schema.Schema{
			"api_key_required": {
				Type: schema.TypeBool,
				Required: false,
			},
			"authorization_scopes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
				Set: schema.HashString,
			},
			"authorization_type": {
				Type: schema.TypeString,
				Required: false,
			},
			"authorizer_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"http_method": {
				Type: schema.TypeString,
				Required: true,
			},
			"integration": {
				Type: schema.TypeList,
				Elem: propertyMethodIntegration(),
				Required: false,
				MaxItems: 1,
			},
			"method_responses": {
				Type: schema.TypeSet,
				Elem: propertyMethodMethodResponse(),
				Required: false,
			},
			"operation_name": {
				Type: schema.TypeString,
				Required: false,
			},
			"request_models": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: false,
			},
			"request_parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeBool},
				Required: false,
			},
			"request_validator_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApiGatewayMethodCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Method", data, meta)
}

func resourceApiGatewayMethodRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Method", data, meta)
}

func resourceApiGatewayMethodUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Method", data, meta)
}

func resourceApiGatewayMethodDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Method", data, meta)
}