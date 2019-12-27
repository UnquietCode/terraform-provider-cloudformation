// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-route.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2RouteType string = "AWS::ApiGatewayV2::Route"

func DatasourceApiGatewayV2Route() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayV2RouteRead,
		
		Schema: map[string]*schema.Schema{
			"target": {
				Type: schema.TypeString,
				Optional: true,
			},
			"route_response_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorizer_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"request_models": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"operation_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"authorization_scopes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"api_key_required": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"route_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"authorization_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"model_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"request_parameters": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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

func datasourceApiGatewayV2RouteRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2RouteType, DatasourceApiGatewayV2Route(), data, meta)
}
