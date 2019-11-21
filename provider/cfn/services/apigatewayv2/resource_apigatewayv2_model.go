// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-model.html

package apigatewayv2

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayV2ModelType string = "AWS::ApiGatewayV2::Model"

func ResourceApiGatewayV2Model() *schema.Resource {
	return &schema.Resource{
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
				Type: schema.TypeString,
				Required: true,
				ValidateFunc: validation.ValidateJsonString,
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
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`[A-Za-z][A-Za-z0-9]+`), "must match pattern [A-Za-z][A-Za-z0-9]+"),
			},
		},
	}
}

func resourceApiGatewayV2ModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayV2ModelType, ResourceApiGatewayV2Model(), data, meta)
}

func resourceApiGatewayV2ModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayV2ModelType, ResourceApiGatewayV2Model(), data, meta)
}

func resourceApiGatewayV2ModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayV2ModelType, ResourceApiGatewayV2Model(), data, meta)
}

func resourceApiGatewayV2ModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayV2ModelType, data, meta)
}

func resourceApiGatewayV2ModelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayV2ModelType, data, meta)
}
