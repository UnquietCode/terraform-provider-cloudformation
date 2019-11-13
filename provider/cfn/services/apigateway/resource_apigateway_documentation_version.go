// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationversion.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayDocumentationVersion() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayDocumentationVersionExists,
		Read:   resourceApiGatewayDocumentationVersionRead,
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
	return plugin.ResourceRead("AWS::ApiGateway::DocumentationVersion", ResourceApiGatewayDocumentationVersion(), data, meta)
}

func resourceApiGatewayDocumentationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::DocumentationVersion", ResourceApiGatewayDocumentationVersion(), data, meta)
}

func resourceApiGatewayDocumentationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::DocumentationVersion", ResourceApiGatewayDocumentationVersion(), data, meta)
}

func resourceApiGatewayDocumentationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::DocumentationVersion", data, meta)
}

func resourceApiGatewayDocumentationVersionCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}