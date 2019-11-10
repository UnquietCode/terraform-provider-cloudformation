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

func ResourceDocumentationVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceDocumentationVersionCreate,
		Read:   resourceDocumentationVersionRead,
		Update: resourceDocumentationVersionUpdate,
		Delete: resourceDocumentationVersionDelete,

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

func resourceDocumentationVersionCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::DocumentationVersion", data, meta)
}

func resourceDocumentationVersionRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::DocumentationVersion", data, meta)
}

func resourceDocumentationVersionUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::DocumentationVersion", data, meta)
}

func resourceDocumentationVersionDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::DocumentationVersion", data, meta)
}