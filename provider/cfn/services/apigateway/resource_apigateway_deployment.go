// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 10-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayDeploymentCreate,
		Read:   resourceApiGatewayDeploymentRead,
		Update: resourceApiGatewayDeploymentUpdate,
		Delete: resourceApiGatewayDeploymentDelete,

		Schema: map[string]*schema.Schema{
			"deployment_canary_settings": {
				Type: schema.TypeList,
				Elem: propertyDeploymentDeploymentCanarySettings(),
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"stage_description": {
				Type: schema.TypeList,
				Elem: propertyDeploymentStageDescription(),
				Optional: true,
				MaxItems: 1,
			},
			"stage_name": {
				Type: schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceApiGatewayDeploymentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::Deployment", data, meta)
}

func resourceApiGatewayDeploymentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::Deployment", data, meta)
}

func resourceApiGatewayDeploymentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::Deployment", data, meta)
}

func resourceApiGatewayDeploymentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::Deployment", data, meta)
}