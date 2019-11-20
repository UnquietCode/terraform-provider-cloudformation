// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 18-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayDeploymentType string = "AWS::ApiGateway::Deployment"

func ResourceApiGatewayDeployment() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayDeploymentRead,
		Create: resourceApiGatewayDeploymentCreate,
		Update: resourceApiGatewayDeploymentUpdate,
		Delete: resourceApiGatewayDeploymentDelete,
		CustomizeDiff: resourceApiGatewayDeploymentCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"deployment_canary_settings": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentDeploymentCanarySettings(),
				Optional: true,
				MaxItems: 1,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"stage_description": {
				Type: schema.TypeSet,
				Elem: propertyDeploymentStageDescription(),
				Optional: true,
				MaxItems: 1,
			},
			"stage_name": {
				Type: schema.TypeString,
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

func resourceApiGatewayDeploymentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayDeploymentType, ResourceApiGatewayDeployment(), data, meta)
}

func resourceApiGatewayDeploymentCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayDeploymentType, ResourceApiGatewayDeployment(), data, meta)
}

func resourceApiGatewayDeploymentUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayDeploymentType, ResourceApiGatewayDeployment(), data, meta)
}

func resourceApiGatewayDeploymentDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayDeploymentType, data, meta)
}

func resourceApiGatewayDeploymentCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayDeploymentType, data, meta)
}
