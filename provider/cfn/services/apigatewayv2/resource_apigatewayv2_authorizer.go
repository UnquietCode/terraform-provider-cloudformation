// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayV2Authorizer() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayV2AuthorizerCreate,
		Read:   resourceApiGatewayV2AuthorizerRead,
		Update: resourceApiGatewayV2AuthorizerUpdate,
		Delete: resourceApiGatewayV2AuthorizerDelete,

		Schema: map[string]*schema.Schema{
			"identity_validation_expression": {
				Type: schema.TypeString,
				Required: false,
			},
			"authorizer_uri": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorizer_credentials_arn": {
				Type: schema.TypeString,
				Required: false,
			},
			"authorizer_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorizer_result_ttl_in_seconds": {
				Type: schema.TypeInt,
				Required: false,
			},
			"identity_source": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceApiGatewayV2AuthorizerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Authorizer", data, meta)
}

func resourceApiGatewayV2AuthorizerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Authorizer", data, meta)
}

func resourceApiGatewayV2AuthorizerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Authorizer", data, meta)
}

func resourceApiGatewayV2AuthorizerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Authorizer", data, meta)
}