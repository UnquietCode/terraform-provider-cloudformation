// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2AuthorizerType string = "AWS::ApiGatewayV2::Authorizer"

func DatasourceApiGatewayV2Authorizer() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayV2AuthorizerRead,
		
		Schema: map[string]*schema.Schema{
			"identity_validation_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_uri": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_credentials_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_type": {
				Type: schema.TypeString,
				Required: true,
			},
			"jwt_configuration": {
				Type: schema.TypeList,
				Elem: propertyAuthorizerJWTConfiguration(),
				Optional: true,
				MaxItems: 1,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceApiGatewayV2AuthorizerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2AuthorizerType, DatasourceApiGatewayV2Authorizer(), data, meta)
}
