// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 15-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html

package apigateway

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayDocumentationPart() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayDocumentationPartExists,
		Read: resourceApiGatewayDocumentationPartRead,
		Create: resourceApiGatewayDocumentationPartCreate,
		Update: resourceApiGatewayDocumentationPartUpdate,
		Delete: resourceApiGatewayDocumentationPartDelete,
		CustomizeDiff: resourceApiGatewayDocumentationPartCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"location": {
				Type: schema.TypeList,
				Elem: propertyDocumentationPartLocation(),
				Required: true,
				MaxItems: 1,
			},
			"properties": {
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

func resourceApiGatewayDocumentationPartExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayDocumentationPartRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGateway::DocumentationPart", ResourceApiGatewayDocumentationPart(), data, meta)
}

func resourceApiGatewayDocumentationPartCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGateway::DocumentationPart", ResourceApiGatewayDocumentationPart(), data, meta)
}

func resourceApiGatewayDocumentationPartUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGateway::DocumentationPart", ResourceApiGatewayDocumentationPart(), data, meta)
}

func resourceApiGatewayDocumentationPartDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGateway::DocumentationPart", data, meta)
}

func resourceApiGatewayDocumentationPartCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff("AWS::ApiGateway::DocumentationPart", data, meta)
}

