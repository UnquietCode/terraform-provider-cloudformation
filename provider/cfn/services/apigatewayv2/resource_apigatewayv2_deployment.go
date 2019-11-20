// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-deployment.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2DeploymentType string = "AWS::ApiGatewayV2::Deployment"

func ResourceApiGatewayV2Deployment() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayV2DeploymentRead,
		Create: resourceApiGatewayV2DeploymentCreate,
		Update: resourceApiGatewayV2DeploymentUpdate,
		Delete: resourceApiGatewayV2DeploymentDelete,
		CustomizeDiff: resourceApiGatewayV2DeploymentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"stage_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"logical_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayV2DeploymentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2DeploymentType, ResourceApiGatewayV2Deployment(), data, meta)
}

func resourceApiGatewayV2DeploymentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayV2DeploymentType, ResourceApiGatewayV2Deployment(), data, meta)
}

func resourceApiGatewayV2DeploymentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayV2DeploymentType, ResourceApiGatewayV2Deployment(), data, meta)
}

func resourceApiGatewayV2DeploymentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayV2DeploymentType, data, meta)
}

func resourceApiGatewayV2DeploymentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayV2DeploymentType, data, meta)
}
