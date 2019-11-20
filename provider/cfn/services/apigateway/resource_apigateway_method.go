// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayMethodType string = "AWS::ApiGateway::Method"

func ResourceApiGatewayMethod() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayMethodRead,
		Create: resourceApiGatewayMethodCreate,
		Update: resourceApiGatewayMethodUpdate,
		Delete: resourceApiGatewayMethodDelete,
		CustomizeDiff: resourceApiGatewayMethodCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"api_key_required": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"authorization_scopes": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"authorization_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"http_method": {
				Type: schema.TypeString,
				Required: true,
			},
			"integration": {
				Type: schema.TypeSet,
				Elem: propertyMethodIntegration(),
				Optional: true,
				MaxItems: 1,
			},
			"method_responses": {
				Type: schema.TypeSet,
				Elem: propertyMethodMethodResponse(),
				Optional: true,
			},
			"operation_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_models": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"request_parameters": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeBool},
				Optional: true,
			},
			"request_validator_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"resource_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayMethodRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayMethodType, ResourceApiGatewayMethod(), data, meta)
}

func resourceApiGatewayMethodCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayMethodType, ResourceApiGatewayMethod(), data, meta)
}

func resourceApiGatewayMethodUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayMethodType, ResourceApiGatewayMethod(), data, meta)
}

func resourceApiGatewayMethodDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayMethodType, data, meta)
}

func resourceApiGatewayMethodCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayMethodType, data, meta)
}
