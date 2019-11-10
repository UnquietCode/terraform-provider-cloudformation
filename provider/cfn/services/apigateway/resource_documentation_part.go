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

func ResourceDocumentationPart() *schema.Resource {
	return &schema.Resource{
		Create: resourceDocumentationPartCreate,
		Read:   resourceDocumentationPartRead,
		Update: resourceDocumentationPartUpdate,
		Delete: resourceDocumentationPartDelete,

		Schema: map[string]*schema.Schema{
			"location": {
				Type: schema.TypeList,
				Elem: propertyDocumentationPartLocation(),
				Required: true,
				ForceNew: true,
				MaxItems: 1,
			},
			"properties": {
				Type: schema.TypeString,
				Required: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceDocumentationPartCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::DocumentationPart", data, meta)
}

func resourceDocumentationPartRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::DocumentationPart", data, meta)
}

func resourceDocumentationPartUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::DocumentationPart", data, meta)
}

func resourceDocumentationPartDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::DocumentationPart", data, meta)
}