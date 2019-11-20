// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayAuthorizerType string = "AWS::ApiGateway::Authorizer"

func ResourceApiGatewayAuthorizer() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayAuthorizerRead,
		Create: resourceApiGatewayAuthorizerCreate,
		Update: resourceApiGatewayAuthorizerUpdate,
		Delete: resourceApiGatewayAuthorizerDelete,
		CustomizeDiff: resourceApiGatewayAuthorizerCustomizeDiff,
		
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

func resourceApiGatewayAuthorizerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayAuthorizerType, ResourceApiGatewayAuthorizer(), data, meta)
}

func resourceApiGatewayAuthorizerCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayAuthorizerType, ResourceApiGatewayAuthorizer(), data, meta)
}

func resourceApiGatewayAuthorizerUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayAuthorizerType, ResourceApiGatewayAuthorizer(), data, meta)
}

func resourceApiGatewayAuthorizerDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayAuthorizerType, data, meta)
}

func resourceApiGatewayAuthorizerCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayAuthorizerType, data, meta)
}
