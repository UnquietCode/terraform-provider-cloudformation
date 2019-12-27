// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-integrationresponse.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2IntegrationResponseType string = "AWS::ApiGatewayV2::IntegrationResponse"

func DatasourceApiGatewayV2IntegrationResponse() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayV2IntegrationResponseRead,
		
		Schema: map[string]*schema.Schema{
			"response_templates": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"template_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"response_parameters": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"content_handling_strategy": {
				Type: schema.TypeString,
				Optional: true,
			},
			"integration_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"integration_response_key": {
				Type: schema.TypeString,
				Required: true,
			},
			"api_id": {
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

func datasourceApiGatewayV2IntegrationResponseRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2IntegrationResponseType, DatasourceApiGatewayV2IntegrationResponse(), data, meta)
}
