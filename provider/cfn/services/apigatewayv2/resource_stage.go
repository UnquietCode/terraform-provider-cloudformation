// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceStage() *schema.Resource {
	return &schema.Resource{
		Create: resourceStageCreate,
		Read:   resourceStageRead,
		Update: resourceStageUpdate,
		Delete: resourceStageDelete,

		Schema: map[string]*schema.Schema{
			"client_certificate_id": {
				Type: schema.TypeString,
				Required: false,
			},
			"deployment_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"access_log_settings": {
				Type: schema.TypeList,
				Elem: propertyStageAccessLogSettings(),
				Required: false,
				MaxItems: 1,
			},
			"route_settings": {
				Type: schema.TypeMap,
				Required: false,
			},
			"stage_name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stage_variables": {
				Type: schema.TypeMap,
				Required: false,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_route_settings": {
				Type: schema.TypeList,
				Elem: propertyStageRouteSettings(),
				Required: false,
				MaxItems: 1,
			},
			"tags": {
				Type: schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceStageCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Stage", data, meta)
}

func resourceStageRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Stage", data, meta)
}

func resourceStageUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Stage", data, meta)
}

func resourceStageDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Stage", data, meta)
}