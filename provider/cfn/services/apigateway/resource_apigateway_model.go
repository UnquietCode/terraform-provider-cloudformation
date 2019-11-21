// This file is generated, and any modifications will be lost when the
// file is next recreated.
//
// Generated on 20-11-2019, using version 0.0 of the cfn terraform provider,
// and version 7.2.0 of the CloudFormation resource specification.
//
// For more information, visit:
//   http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html

package apigateway

import (
	"regexp"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/unquietcode/terraform-cfn-provider/plugin"
)

const apiGatewayModelType string = "AWS::ApiGateway::Model"

func ResourceApiGatewayModel() *schema.Resource {
	return &schema.Resource{
		Read: resourceApiGatewayModelRead,
		Create: resourceApiGatewayModelCreate,
		Update: resourceApiGatewayModelUpdate,
		Delete: resourceApiGatewayModelDelete,
		CustomizeDiff: resourceApiGatewayModelCustomizeDiff,
		
		Schema: map[string]*schema.Schema{
			"content_type": {
				Type: schema.TypeString,
				Optional: true,
			},
			"description": {
				Type: schema.TypeString,
				Optional: true,
			},
			"name": {
				Type: schema.TypeString,
				Optional: true,
			},
			"rest_api_id": {
				Type: schema.TypeString,
				Required: true,
			},
			"schema": {
				Type: schema.TypeString,
				Optional: true,
				ValidateFunc: validation.ValidateJsonString,
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

func resourceApiGatewayModelRead(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceRead(apiGatewayModelType, ResourceApiGatewayModel(), data, meta)
}

func resourceApiGatewayModelCreate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceCreate(apiGatewayModelType, ResourceApiGatewayModel(), data, meta)
}

func resourceApiGatewayModelUpdate(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceUpdate(apiGatewayModelType, ResourceApiGatewayModel(), data, meta)
}

func resourceApiGatewayModelDelete(data *schema.ResourceData, meta interface{}) error {
	return plugin.ResourceDelete(apiGatewayModelType, data, meta)
}

func resourceApiGatewayModelCustomizeDiff(data *schema.ResourceDiff, meta interface{}) error {
	return plugin.ResourceCustomizeDiff(apiGatewayModelType, data, meta)
}
