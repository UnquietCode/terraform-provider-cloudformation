// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2StageType string = "AWS::ApiGatewayV2::Stage"

var apiGatewayV2StageProperties map[string]string = map[string]string{
	"client_certificate_id": "ClientCertificateId",
	"deployment_id": "DeploymentId",
	"description": "Description",
	"access_log_settings": "AccessLogSettings",
	"route_settings": "RouteSettings",
	"stage_name": "StageName",
	"stage_variables": "StageVariables",
	"api_id": "ApiId",
	"default_route_settings": "DefaultRouteSettings",
	"tags": "Tags",
}

func ResourceApiGatewayV2Stage() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayV2StageExists,
		Read: resourceApiGatewayV2StageRead,
		Create: resourceApiGatewayV2StageCreate,
		Update: resourceApiGatewayV2StageUpdate,
		Delete: resourceApiGatewayV2StageDelete,
		CustomizeDiff: resourceApiGatewayV2StageCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"client_certificate_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"deployment_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"access_log_settings": {
				Type: schema.TypeList,
				Elem: propertyStageAccessLogSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"route_settings": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"stage_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stage_variables": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_route_settings": {
				Type: schema.TypeList,
				Elem: propertyStageRouteSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayV2StageExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayV2StageRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2StageType, ResourceApiGatewayV2Stage(), data, meta)
}

func resourceApiGatewayV2StageCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayV2StageType, ResourceApiGatewayV2Stage(), data, apiGatewayV2StageProperties, meta)
}

func resourceApiGatewayV2StageUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayV2StageType, ResourceApiGatewayV2Stage(), data, apiGatewayV2StageProperties, meta)
}

func resourceApiGatewayV2StageDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayV2StageType, data, meta)
}

func resourceApiGatewayV2StageCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayV2StageType, data, meta)
}
