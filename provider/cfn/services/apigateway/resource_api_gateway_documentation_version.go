// This file is generated, and any modifications will be lost
// when the file is next recreated.
//
// Generated on 07-11-2019, using version 0.0 of the cfn
// terraform provider, and version 7.2.0 of the CloudFormation
// resource specification.

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayDocumentationVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceApiGatewayDocumentationVersionCreate,
		Read:   resourceApiGatewayDocumentationVersionRead,
		Update: resourceApiGatewayDocumentationVersionUpdate,
		Delete: resourceApiGatewayDocumentationVersionDelete,

		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Required: false,
			},
			"documentation_version": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceApiGatewayDocumentationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::DocumentationVersion", data, meta)
}

func resourceApiGatewayDocumentationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::DocumentationVersion", data, meta)
}

func resourceApiGatewayDocumentationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::DocumentationVersion", data, meta)
}

func resourceApiGatewayDocumentationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::DocumentationVersion", data, meta)
}