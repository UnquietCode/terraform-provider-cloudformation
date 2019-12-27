// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-api.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2ApiType string = "AWS::ApiGatewayV2::Api"

func DatasourceApiGatewayV2Api() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayV2ApiRead,
		
		Schema: map[string]*schema.Schema{
			"route_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
			},
			"body_s3_location": {
				Type: schema.TypeList,
				Elem: propertyApiBodyS3Location(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"base_path": {
				Type: schema.TypeString,
				Optional: true,
			},
			"fail_on_warnings": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"disable_schema_validation": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"target": {
				Type: schema.TypeString,
				Optional: true,
			},
			"credentials_arn": {
				Type: schema.TypeString,
				Optional: true,
			},
			"cors_configuration": {
				Type: schema.TypeList,
				Elem: propertyApiCors(),
				Optional: true,
				MaxItems: 1,
			},
			"version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"protocol_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"route_key": {
				Type: schema.TypeString,
				Optional: true,
			},
			"body": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"tags": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"api_key_selection_expression": {
				Type: schema.TypeString,
				Optional: true,
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

func datasourceApiGatewayV2ApiRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2ApiType, DatasourceApiGatewayV2Api(), data, meta)
}
