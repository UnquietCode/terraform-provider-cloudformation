// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 09-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceDeploymentCreate,
		Read:   resourceDeploymentRead,
		Update: resourceDeploymentUpdate,
		Delete: resourceDeploymentDelete,

		Schema: map[string]*schema.Schema{
			"deployment_canary_settings": {
				Type: schema.TypeList,
				Elem: propertyDeploymentDeploymentCanarySettings(),
				Required: false,
				ForceNew: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stage_description": {
				Type: schema.TypeList,
				Elem: propertyDeploymentStageDescription(),
				Required: false,
				MaxItems: 1,
			},
			"stage_name": {
				Type: schema.TypeString,
				Required: false,
			},
		},
	}
}

func resourceDeploymentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Deployment", data, meta)
}

func resourceDeploymentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Deployment", data, meta)
}

func resourceDeploymentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Deployment", data, meta)
}

func resourceDeploymentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Deployment", data, meta)
}