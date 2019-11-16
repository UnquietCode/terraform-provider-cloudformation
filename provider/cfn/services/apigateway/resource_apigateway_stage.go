// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/cfn/misc"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayStageType string = "AWS::ApiGateway::Stage"

var apiGatewayStageProperties map[string]string = map[string]string{
	"access_log_setting": "AccessLogSetting",
	"cache_cluster_enabled": "CacheClusterEnabled",
	"cache_cluster_size": "CacheClusterSize",
	"canary_setting": "CanarySetting",
	"client_certificate_id": "ClientCertificateId",
	"deployment_id": "DeploymentId",
	"description": "Description",
	"documentation_version": "DocumentationVersion",
	"method_settings": "MethodSettings",
	"rest_api_id": "RestApiId",
	"stage_name": "StageName",
	"tags": "Tags",
	"tracing_enabled": "TracingEnabled",
	"variables": "Variables",
}

func ResourceApiGatewayStage() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayStageExists,
		Read: resourceApiGatewayStageRead,
		Create: resourceApiGatewayStageCreate,
		Update: resourceApiGatewayStageUpdate,
		Delete: resourceApiGatewayStageDelete,
		CustomizeDiff: resourceApiGatewayStageCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"access_log_setting": {
				Type: schema.TypeList,
				Elem: propertyStageAccessLogSetting(),
				Optional: true,
				MaxItems: 1,
			},
			"cache_cluster_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"cache_cluster_size": {
				Type: schema.TypeString,
				Optional: true,
			},
			"canary_setting": {
				Type: schema.TypeList,
				Elem: propertyStageCanarySetting(),
				Optional: true,
				MaxItems: 1,
			},
			"client_certificate_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"deployment_id": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"documentation_version": {
				Type: schema.TypeString,
				Optional: true,
			},
			"method_settings": {
				Type: schema.TypeSet,
				Elem: propertyStageMethodSetting(),
				Optional: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"stage_name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"tags": misc.PropertyTags(),
			"tracing_enabled": {
				Type: schema.TypeBool,
				Optional: true,
			},
			"variables": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{Type: schema.TypeString},
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

func resourceApiGatewayStageExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayStageRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayStageType, ResourceApiGatewayStage(), data, meta)
}

func resourceApiGatewayStageCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayStageType, ResourceApiGatewayStage(), data, apiGatewayStageProperties, meta)
}

func resourceApiGatewayStageUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayStageType, ResourceApiGatewayStage(), data, apiGatewayStageProperties, meta)
}

func resourceApiGatewayStageDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayStageType, data, meta)
}

func resourceApiGatewayStageCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayStageType, data, meta)
}
