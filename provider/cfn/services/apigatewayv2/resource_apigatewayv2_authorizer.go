// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2AuthorizerType string = "AWS::ApiGatewayV2::Authorizer"

var apiGatewayV2AuthorizerProperties map[string]string = map[string]string{
	"identity_validation_expression": "IdentityValidationExpression",
	"authorizer_uri": "AuthorizerUri",
	"authorizer_credentials_arn": "AuthorizerCredentialsArn",
	"authorizer_type": "AuthorizerType",
	"authorizer_result_ttl_in_seconds": "AuthorizerResultTtlInSeconds",
	"identity_source": "IdentitySource",
	"api_id": "ApiId",
	"name": "Name",
}

func ResourceApiGatewayV2Authorizer() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayV2AuthorizerExists,
		Read: resourceApiGatewayV2AuthorizerRead,
		Create: resourceApiGatewayV2AuthorizerCreate,
		Update: resourceApiGatewayV2AuthorizerUpdate,
		Delete: resourceApiGatewayV2AuthorizerDelete,
		CustomizeDiff: resourceApiGatewayV2AuthorizerCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"identity_validation_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_uri": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorizer_credentials_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorizer_result_ttl_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"identity_source": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
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

func resourceApiGatewayV2AuthorizerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayV2AuthorizerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2AuthorizerType, ResourceApiGatewayV2Authorizer(), data, meta)
}

func resourceApiGatewayV2AuthorizerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayV2AuthorizerType, ResourceApiGatewayV2Authorizer(), data, apiGatewayV2AuthorizerProperties, meta)
}

func resourceApiGatewayV2AuthorizerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayV2AuthorizerType, ResourceApiGatewayV2Authorizer(), data, apiGatewayV2AuthorizerProperties, meta)
}

func resourceApiGatewayV2AuthorizerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayV2AuthorizerType, data, meta)
}

func resourceApiGatewayV2AuthorizerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayV2AuthorizerType, data, meta)
}
