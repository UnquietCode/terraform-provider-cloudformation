// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 16-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayDocumentationVersionType string = "AWS::ApiGateway::DocumentationVersion"

var apiGatewayDocumentationVersionProperties map[string]string = map[string]string{
	"description": "Description",
	"documentation_version": "DocumentationVersion",
	"rest_api_id": "RestApiId",
}

func ResourceApiGatewayDocumentationVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayDocumentationVersionExists,
		Read: resourceApiGatewayDocumentationVersionRead,
		Create: resourceApiGatewayDocumentationVersionCreate,
		Update: resourceApiGatewayDocumentationVersionUpdate,
		Delete: resourceApiGatewayDocumentationVersionDelete,
		CustomizeDiff: resourceApiGatewayDocumentationVersionCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"documentation_version": {
				Type: schema.TypeString,
				Required: true,
			},
			"rest_api_id": {
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

func resourceApiGatewayDocumentationVersionExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayDocumentationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayDocumentationVersionType, ResourceApiGatewayDocumentationVersion(), data, meta)
}

func resourceApiGatewayDocumentationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayDocumentationVersionType, ResourceApiGatewayDocumentationVersion(), data, apiGatewayDocumentationVersionProperties, meta)
}

func resourceApiGatewayDocumentationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayDocumentationVersionType, ResourceApiGatewayDocumentationVersion(), data, apiGatewayDocumentationVersionProperties, meta)
}

func resourceApiGatewayDocumentationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayDocumentationVersionType, data, meta)
}

func resourceApiGatewayDocumentationVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayDocumentationVersionType, data, meta)
}
