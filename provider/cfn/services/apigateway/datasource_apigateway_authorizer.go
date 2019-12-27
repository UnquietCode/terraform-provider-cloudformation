// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayAuthorizerType string = "AWS::ApiGateway::Authorizer"

func DatasourceApiGatewayAuthorizer() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayAuthorizerRead,
		
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceApiGatewayAuthorizerRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayAuthorizerType, DatasourceApiGatewayAuthorizer(), data, meta)
}
