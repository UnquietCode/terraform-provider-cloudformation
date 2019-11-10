// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayV2Stage() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayV2StageCreate,
		Read:   resourceApiGatewayV2StageRead,
		Update: resourceApiGatewayV2StageUpdate,
		Delete: resourceApiGatewayV2StageDelete,

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
				ForceNew: true,
			},
			"stage_variables": {
				Type: schema.TypeMap,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
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
		},
	}
}

func resourceApiGatewayV2StageCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Stage", ResourceApiGatewayV2Stage(), data, meta)
}

func resourceApiGatewayV2StageRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Stage", ResourceApiGatewayV2Stage(), data, meta)
}

func resourceApiGatewayV2StageUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Stage", ResourceApiGatewayV2Stage(), data, meta)
}

func resourceApiGatewayV2StageDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Stage", data, meta)
}