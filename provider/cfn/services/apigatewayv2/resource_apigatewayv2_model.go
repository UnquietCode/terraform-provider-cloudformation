// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 13-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html

package apigatewayv2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

func ResourceApiGatewayV2Model() *schema.Resource {
	return &schema.Resource{
		Exists: resourceApiGatewayV2ModelExists,
		Read: resourceApiGatewayV2ModelRead,
		Create: resourceApiGatewayV2ModelCreate,
		Update: resourceApiGatewayV2ModelUpdate,
		Delete: resourceApiGatewayV2ModelDelete,
		CustomizeDiff: resourceApiGatewayV2ModelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"content_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"schema": {
				Type: schema.TypeMap,
				Required: true,
			},
			"api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"name": {
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

func resourceApiGatewayV2ModelExists(data *schema.ResourceData, meta interface{}) (bool, error) {
	return plugin.ResourceExists(data, meta)
}

func resourceApiGatewayV2ModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead("AWS::ApiGatewayV2::Model", ResourceApiGatewayV2Model(), data, meta)
}

func resourceApiGatewayV2ModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate("AWS::ApiGatewayV2::Model", ResourceApiGatewayV2Model(), data, meta)
}

func resourceApiGatewayV2ModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate("AWS::ApiGatewayV2::Model", ResourceApiGatewayV2Model(), data, meta)
}

func resourceApiGatewayV2ModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete("AWS::ApiGatewayV2::Model", data, meta)
}

func resourceApiGatewayV2ModelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(data, meta)
}
