// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayMethodType string = "AWS::ApiGateway::Method"

func DatasourceApiGatewayMethod() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayMethodRead,
		
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
				Type: schema.TypeList,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceApiGatewayMethodRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayMethodType, DatasourceApiGatewayMethod(), data, meta)
}
