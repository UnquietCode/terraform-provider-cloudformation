// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayAuthorizer() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayAuthorizerExists,
		Read:   resourceApiGatewayAuthorizerRead,
		Create: resourceApiGatewayAuthorizerCreate,
		Update: resourceApiGatewayAuthorizerUpdate,
		Delete: resourceApiGatewayAuthorizerDelete,
		
		Schema: map[string]*schema.Schema{
			"auth_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_credentials": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_result_ttl_in_seconds": {
				Type: schema.TypeInt,
				Optional: true,
			},
			"authorizer_uri": {
				Type: schema.TypeString,
				Optional: true,
			},
			"identity_source": {
				Type: schema.TypeString,
				Optional: true,
			},
			"identity_validation_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"provider_ar_ns": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Set: schema.HashString,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"type": {
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

func resourceApiGatewayAuthorizerExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayAuthorizerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Authorizer", ResourceApiGatewayAuthorizer(), data, meta)
}

func resourceApiGatewayAuthorizerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Authorizer", ResourceApiGatewayAuthorizer(), data, meta)
}

func resourceApiGatewayAuthorizerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Authorizer", ResourceApiGatewayAuthorizer(), data, meta)
}

func resourceApiGatewayAuthorizerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Authorizer", data, meta)
}