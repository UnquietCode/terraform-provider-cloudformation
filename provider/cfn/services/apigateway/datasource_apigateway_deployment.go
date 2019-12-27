// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 26-12-2019, using version 0.0 of the cfn terraform provider,
// and version 10.1.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayDeploymentType string = "AWS::ApiGateway::Deployment"

func DatasourceApiGatewayDeployment() *schema.Resource {
	return &schema.Resource{
		Read: datasourceApiGatewayDeploymentRead,
		
		Schema: map[string]*schema.Schema{
			"deployment_canary_settings": {
				Type: schema.TypeList,
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
				Type: schema.TypeList,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
			"json": {
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceApiGatewayDeploymentRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayDeploymentType, DatasourceApiGatewayDeployment(), data, meta)
}
