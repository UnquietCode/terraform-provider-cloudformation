// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2StageType string = "AWS::ApiGatewayV2::Stage"

func ResourceApiGatewayV2Stage() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeSet,
				Elem: propertyStageAccessLogSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"route_settings": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"stage_name": {
				Type: schema.TypeString,
				Required: true,
			},
			"stage_variables": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"default_route_settings": {
				Type: schema.TypeSet,
				Elem: propertyStageRouteSettings(),
				Optional: true,
				MaxItems: 1,
			},
			"tags": {
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
		},
	}
}

func resourceApiGatewayV2StageRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2StageType, ResourceApiGatewayV2Stage(), data, meta)
}

func resourceApiGatewayV2StageCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayV2StageType, ResourceApiGatewayV2Stage(), data, meta)
}

func resourceApiGatewayV2StageUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayV2StageType, ResourceApiGatewayV2Stage(), data, meta)
}

func resourceApiGatewayV2StageDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayV2StageType, data, meta)
}

func resourceApiGatewayV2StageCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayV2StageType, data, meta)
}
